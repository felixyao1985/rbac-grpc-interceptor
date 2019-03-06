package unit

import (
	"rbac-grpc-interceptor/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

const (
	SecretKey = "felix jwt demo"
)

type Token struct {
	Token string `json:"token"`
}
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//生成token
func GenerateToken(username string,pwd string) (Token,error){
	/*
		加入登陆判断，并获得相应权限
	*/
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = username
	claims["role"] = [...]string{"Member","Admin"}
	claims["pwd"] = pwd
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println("Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}

	return response,err
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string) (token *jwt.Token, err error) {
	//var token *jwt.Token
	kv := strings.Split(tokenSrt, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil, nil
	}
	tokenString := kv[1]

	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return
}

//JWT 鉴权中间件--高阶函数
func ValidateMiddleware() grpc.UnaryClientInterceptor{

	interceptor := grpc_middleware.ChainUnaryClient(TokenUnaryClient)

	return interceptor
}

//验证Token
func TokenUnaryClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) ( error) {

	authorization,_ := extractClientHeader(ctx,"authorization")
	fmt.Println(authorization)
	err := invoker(ctx, method, req, reply, cc, opts...)

	fmt.Println("server_role_manage Client 调用了interceptro")

	return err
}

func TokenUnaryServerChain(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	authorization,_ := extractServerHeader(ctx,"authorization")

	token, err := ParseToken(authorization)

	if(err ==nil && token.Valid) {
		return handler(ctx, req)
	}else{
		//验证不通过
		return handler(ctx, req)
	}

}

//验证权限

func RBACUnaryServerChain(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	t := test.NewRBACClient(conn)

	FullMethod := strings.Split(info.FullMethod, "/")
	authorization,_ := extractServerHeader(ctx,"authorization")

	token, err := ParseToken(authorization)
	claims, _ := token.Claims.(jwt.MapClaims)

	role,_ := claims["role"]
	roles := reflect.ValueOf(role)
	Sub := []string{}
	for i := 0; i < roles.Len(); i++ {
		st := roles.Index(i)
		Sub = append(Sub,st.Interface().(string))
	}
	fmt.Println(Sub)
	tr, err := t.Enforce(context.Background(),&test.EnforceRep{
		Sub:Sub,
		Obj:FullMethod[1],
		Act:FullMethod[2],
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("服务端响应: %s", tr.Code)
	if(tr.Code==0) {
		return nil, grpc.Errorf(codes.Unauthenticated, "Permission denied")
	}else{
		return handler(ctx, req)
	}

}


func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func Login(w http.ResponseWriter, r *http.Request) {
	URLValues := r.URL.Query()

	response := map[string]string{}
	response["username"] = URLValues.Get("username")
	Token,_ := GenerateToken(URLValues.Get("username"),URLValues.Get("username"))
	response["token"] = Token.Token

	JsonResponse(response,w)
}

func extractClientHeader(ctx context.Context, header string) (string, error) {
	md,_, ok := metadata.FromOutgoingContextRaw(ctx)

	if !ok {
		return "", status.Error(codes.Unauthenticated, "no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	if len(authHeaders) != 1 {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeaders[0], nil
}

func extractServerHeader(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", status.Error(codes.Unauthenticated, "no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	if len(authHeaders) != 1 {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeaders[0], nil
}