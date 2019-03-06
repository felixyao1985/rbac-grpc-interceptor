package authentication

import (
	"strings"
)

type ModAct struct{
	Sub string
	Obj string
	Act string
}
type authentication struct{
	model map[string][]string
	permitacts map[string]bool
}

func New() *authentication {

	return &authentication{}
}

func (e *authentication) LoadModel(list []*ModAct) {
	permitacts := map[string]bool{}
	model := map[string][]string{}
	for i := 0; i < len(list); i++ {
		item := list[i]
		if &item != nil {

			if model[item.Sub] != nil && len(model[item.Sub]) == 0 {
				model[item.Sub] = []string{}
			}
			Sub := item.Sub
			Obj := item.Obj
			Act := item.Act
			if(Sub == "*" || strings.ToUpper(Sub)=="ALL") {
				Sub = "SubALL"
			}
			if(Obj == "*" || strings.ToUpper(Obj)=="ALL") {
				Obj = "ObjALL"
			}
			if(Act == "*" || strings.ToUpper(Act)=="ALL") {
				Act = "ActALL"
			}

			permitact := Sub + "_" +Obj + "_"+ Act

			permitacts[permitact] = true
			model[Sub] = append(model[Sub],Obj)
		}
	}

	e.permitacts = permitacts
	e.model = model
}

func (e *authentication) GetModel() map[string][]string {
	return e.model
}

func (e *authentication)IsPpermited(Sub string,Obj string,Act string) bool {

	permitact := Sub + "_" +Obj + "_"+ Act

	//限定角色 模块 操作
	if(e.permitacts[permitact]) {
		return true
	}

	return false
}

func (e *authentication) Enforce(Sub string,Obj string,Act string) bool {

	if (e.IsPpermited(Sub ,Obj ,Act )) {
		return true
	}

	//一个模块 给所有角色 开所有权限
	if (e.IsPpermited("SubALL" ,Obj ,"ActALL" )) {
		return true
	}

	//指定角色 模块 开启所有权限
	if (e.IsPpermited(Sub ,Obj ,"ActALL" )) {
		return true
	}

	//指定角色 所有模块 开启所有权限
	if (e.IsPpermited(Sub ,"ActALL" ,"ActALL" )) {
		return true
	}

	//指定角色 开启所有模块 指定权限
	if (e.IsPpermited(Sub ,"ObjALL" ,Act )) {
		return true
	}

	//所有角色 指定模块 指定权限
	if (e.IsPpermited("SubALL" ,Obj ,Act )) {
		return true
	}

	return false
}