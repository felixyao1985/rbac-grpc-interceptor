package interceptor

import (
	"strings"
)

type Permit struct {
	Role    string
	Method  string
	Pattern string
	Module  string
}
type authentication struct {
	modules    map[string][]string
	permitacts map[string]bool
}

func New() *authentication {

	return &authentication{}
}

func (e *authentication) LoadModel(list []*Permit) {
	permitacts := map[string]bool{}
	modules := map[string][]string{}
	for i := 0; i < len(list); i++ {
		item := list[i]
		if &item != nil {

			Sub := item.Role
			Obj := item.Method
			Act := item.Pattern
			if Sub == "*" || strings.ToUpper(Sub) == "ALL" {
				Sub = "SubALL"
			}
			if Obj == "*" || strings.ToUpper(Obj) == "ALL" {
				Obj = "ObjALL"
			}
			if Act == "*" || strings.ToUpper(Act) == "ALL" {
				Act = "ActALL"
			}

			permitact := Sub + "_" + Obj + "_" + Act
			permitacts[permitact] = true

			if modules[item.Role] != nil && len(modules[item.Role]) == 0 {
				modules[item.Role] = []string{}
			}
			modules[Sub] = append(modules[Sub], item.Module)
		}
	}

	e.permitacts = permitacts
	e.modules = modules
}

func (e *authentication) GetModel() map[string][]string {
	return e.modules
}

func (e *authentication) IsPpermited(Sub string, Obj string, Act string) bool {

	permitact := Sub + "_" + Obj + "_" + Act
	if e.permitacts[permitact] {
		return true
	}

	return false
}

func (e *authentication) Enforce(Sub string, Obj string, Act string) bool {

	if e.IsPpermited(Sub, Obj, Act) {
		return true
	}

	if e.IsPpermited("SubALL", Obj, "ActALL") {
		return true
	}

	if e.IsPpermited(Sub, Obj, "ActALL") {
		return true
	}

	if e.IsPpermited(Sub, "ActALL", "ActALL") {
		return true
	}

	if e.IsPpermited(Sub, "ObjALL", Act) {
		return true
	}

	if e.IsPpermited("SubALL", Obj, Act) {
		return true
	}

	return false
}
