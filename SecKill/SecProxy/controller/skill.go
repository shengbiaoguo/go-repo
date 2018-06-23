package controller

import "github.com/astaxie/beego"

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	p.Data["json"] = "秒杀"
	p.ServeJSON()
}

func (p *SkillController) SecInfo() {
	p.Data["json"] = "秒杀信息"
	p.ServeJSON()
}
