// 项目简介 oj网站
测试样例： 发送邮件->注册->登录-> 问题创建（admin） -> 问题修改-> 问题列表->问题详情 ->提交列表（代码判断）
// 技术点：表关联，md5，JWT,  邮箱验证码，redis,  go-uuid，排序查询，鉴权(普通人员，管理员 通过自定义中间件)
get /get_problem_list
category_identity:  Category_1
