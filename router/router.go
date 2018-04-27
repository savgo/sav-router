package router

import (
  "strconv"
  "github.com/savgo/sav-util/strcase"
  "github.com/jetiny/route"
)

type Method uint8

const (
  GET = iota
  POST
  PUT
  DELETE
  PATCH
  ANY
)

var Methods = map[Method]string{
  GET: "GET",
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
  PATCH: "PATCH",
  ANY: "ANY",
}

type RouterOptions struct {
  Prefix string
  Sensitive bool
  Method Method
  CaseType strcase.CaseType
}

type ModalRoute struct {
  Name string
  Path string
  Opts map[string]interface{}
  childs map[Method][]*ActionRoute
  route route.Route
}

type ActionRoute struct {
  Name string
}

type Router struct {
  opts RouterOptions
  modalMap map[string]*ModalRoute
  modalRoutes []*ModalRoute
  actionRoutes []*ActionRoute
  absoluteRoutes []*ActionRoute
}

func (router * Router) createModalRoute(opts map[string]interface{}) {
  name := opts["name"].(string)
  modalRoute := &ModalRoute{
    Name: strcase.Pascal(name),
    Path: strcase.Convert(name, router.opts.CaseType),
    Opts: opts,
  }
  // 处理路径
  _, ok := opts["path"]
  if ok {
    switch opts["path"].(type) {
      case string:
        modalRoute.Path = opts["path"].(string)
    }
  }
//@TODO $route['path'] = $this->normalPath('/' . $path);
  // 生成路由
  modalRoute.route = route.Parse(modalRoute.Path, &route.ParseOption{
    End: false,
    Sensitive: router.opts.Sensitive,
  })
  // 添加到modalMap
  router.modalMap[name] = modalRoute
  _, ok = opts["id"]
  if ok {
    switch v := opts["id"].(type) {
      case string:
        router.modalMap[opts["id"].(string)] = modalRoute
      case int:
        router.modalMap[strconv.FormatInt(int64(v), 10)] = modalRoute
    }
  }
  router.modalRoutes = append(router.modalRoutes, modalRoute)
// if (isset($opts['routes'])) {
//     foreach ($opts['routes'] as $key => $it) {
//         if (!is_numeric($key)) {
//             $it['name'] = $key;
//         }
//         $it['modal'] = $opts['name'];
//         $this->build($it);
//     }
// }
// return $route;
}
