# ENTERING FILE: add_class.ts
var("hi") {
  set("hello")
}
attribute("class") {
  value() {
    append(" ")
    append(var("hi"))
  }
  replace("h", "H")
}
# LEAVING FILE: add_class.ts