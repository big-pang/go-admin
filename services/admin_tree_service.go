package services

import (
	"github.com/beego/beego/v2/client/orm"
	"go-admin/models"
	"go-admin/utils"
	"sort"
	"strconv"
	"strings"
)

// AdminTreeService struct
type AdminTreeService struct {
	Ret   string
	Html  string
	Array map[int]orm.Params
	Text  map[string]interface{}
}

var (
	icon  []string = []string{"│", "├", "└"}
	space string   = "&nbsp;&nbsp;"
)

// GetLeftMenu 获取左侧菜单
func (adminTreeService *AdminTreeService) GetLeftMenu(requestUrl string, user models.AdminUser) string {
	menu := user.GetShowMenu()
	maxLevel := 0
	currentID := 1
	parentIds := []int{0}

	for _, v := range menu {
		if v["Url"].(string) == requestUrl {
			parentIds = adminTreeService.getMenuParent(menu, int(v["Id"].(int64)), []int{})
			currentID = int(v["Id"].(int64))
		}
	}

	if len(parentIds) == 0 {
		parentIds = []int{0}
	}

	for k, v := range menu {
		menu[k]["Url"] = "/" + v["Url"].(string)
		tempLevel := adminTreeService.GetLevel(int(v["Id"].(int64)), menu, 0)
		menu[k]["Level"] = tempLevel
		if maxLevel <= tempLevel {
			maxLevel = tempLevel
		}
	}

	adminTreeService.initTree(menu)

	textBaseOne := "<li class='treeview"
	textHover := " active"
	textBaseTwo := `'><a href='javascript:void(0);'>
	<i class='fa $icon'></i>
	<span>
	$name
	</span>
	<span class='pull-right-container'><i class='fa fa-angle-left pull-right'></i></span>
	</a><ul class='treeview-menu`
	textOpen := " menu-open"
	textBaseThree := "'>"

	textBaseFour := "<li"
	textHoverLi := " class='active'"
	textBaseFive := `>
	<a href='$url'>
	<i class='fa $icon'></i>
	<span>$name</span>
	</a>
	</li>`

	text0 := textBaseOne + textBaseTwo + textBaseThree
	text1 := textBaseOne + textHover + textBaseTwo + textOpen + textBaseThree
	text2 := "</ul></li>"
	textCurrent := textBaseFour + textHoverLi + textBaseFive
	textOther := textBaseFour + textBaseFive

	adminTreeService.Text = make(map[string]interface{})
	for i := 0; i <= maxLevel; i++ {
		adminTreeService.Text[strconv.Itoa(i)] = []string{text0, text1, text2}
	}
	adminTreeService.Text["current"] = textCurrent
	adminTreeService.Text["other"] = textOther

	return adminTreeService.getAuthTree(0, currentID, parentIds)

}

// getMenuParent 获取父级菜单
func (adminTreeService *AdminTreeService) getMenuParent(menu map[int]orm.Params, myID int, parentIds []int) []int {
	for _, a := range menu {
		if int(a["Id"].(int64)) == myID && int(a["ParentId"].(int64)) != 0 {
			parentIds = append(parentIds, int(a["ParentId"].(int64)))
			parentIds = adminTreeService.getMenuParent(menu, int(a["ParentId"].(int64)), parentIds)
		}
	}
	if len(parentIds) > 0 {
		return parentIds
	}
	return []int{}
}

// GetLevel 递归获取级别
func (adminTreeService *AdminTreeService) GetLevel(id int, menu map[int]orm.Params, i int) int {
	parentID := 0
	v, ok := menu[id]["ParentId"].(int64)
	if ok {
		parentID = int(v)
	} else {
		v1, ok := menu[id]["ParentId"].(int)
		if ok {
			parentID = v1
		}
	}

	if (parentID == 0) || id == parentID {
		return i
	}
	i++
	return adminTreeService.GetLevel(parentID, menu, i)
}

// initTree 初始化树
func (adminTreeService *AdminTreeService) initTree(menu map[int]orm.Params) {
	adminTreeService.Array = make(map[int]orm.Params)
	adminTreeService.Array = menu
	adminTreeService.Ret = ""
	adminTreeService.Html = ""
}

// getAuthTree 获取后台左侧菜单
func (adminTreeService *AdminTreeService) getAuthTree(myId int, currentID int, parentIds []int) string {
	nStr := ""
	child := adminTreeService.getChild(myId)
	if len(child) > 0 {
		menu := make(map[string]interface{})
		//取key最小的一个，防止for range随机取，导致每次菜单顺序不同
		sortID := 99999
		for k := range child {
			if k < sortID {
				sortID = k
			}
		}
		menu = child[sortID]

		//获取当前等级的html
		var textHTMLInterface interface{}
		if adminTreeService.Text[strconv.Itoa(menu["Level"].(int))] != "" {
			//[]string类型
			textHTMLInterface = adminTreeService.Text[strconv.Itoa(menu["Level"].(int))]
		} else {
			//string类型
			textHTMLInterface = adminTreeService.Text["other"]
		}

		//child排序，防止菜单位置每次都不同
		var childKeys []int
		for k := range child {
			childKeys = append(childKeys, k)
		}
		sort.Ints(childKeys)
		for _, key := range childKeys {
			k := key
			v := child[key]

			if len(adminTreeService.getChild(k)) > 0 {
				textHTMLArr := textHTMLInterface.([]string)
				if utils.InArrayForInt(parentIds, k) {
					nStr = adminTreeService.strReplace(textHTMLArr[1], v)
					adminTreeService.Html += nStr
				} else {
					nStr = adminTreeService.strReplace(textHTMLArr[0], v)
					adminTreeService.Html += nStr
				}
				adminTreeService.getAuthTree(k, currentID, parentIds)
				nStr = adminTreeService.strReplace(textHTMLArr[2], v)
				adminTreeService.Html += nStr
			} else if k == currentID {
				a := adminTreeService.Text["current"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			} else {
				a := adminTreeService.Text["other"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			}
		}
	}
	return adminTreeService.Html
}

// getChild 得到子级数组
func (adminTreeService *AdminTreeService) getChild(pid int) map[int]map[string]interface{} {
	result := make(map[int]map[string]interface{})
	for k, v := range adminTreeService.Array {
		parentID, ok := v["ParentId"].(int64)
		var parentIDInt int
		if ok {
			parentIDInt = int(parentID)
		} else {
			parentIDInt = v["ParentId"].(int)
		}

		if parentIDInt == pid {
			result[k] = v
		}
	}
	return result
}

// strReplace 替换字符串
func (adminTreeService *AdminTreeService) strReplace(str string, m map[string]interface{}) string {
	str = strings.ReplaceAll(str, "$icon", m["Icon"].(string))
	str = strings.ReplaceAll(str, "$name", m["Name"].(string))
	str = strings.ReplaceAll(str, "$url", m["Url"].(string))
	return str
}

// GetTree 得到树型结构
//myId 表示获得这个ID下的所有子级
//str 生成树型结构的基本代码，例如："<option value=\$id \$selected>\$spacer\$name</option>"
//sid 被选中的ID，比如在做树型下拉框的时候需要用到.
func (adminTreeService *AdminTreeService) GetTree(myId int, str string, sid int, adds string, strGroup string) string {
	number := 1
	child := adminTreeService.getChild(myId)

	if len(child) > 0 {
		total := len(child)

		//child排序
		var ids []int
		for id := range child {
			ids = append(ids, id)
		}
		sort.Ints(ids)

		for _, id := range ids {
			value := child[id]

			j := ""
			k := ""
			if number == total {
				j += icon[2]
			} else {
				j += icon[1]
				if adds != "" {
					k = icon[0]
				} else {
					k = ""
				}
			}

			spacer := ""
			if adds != "" {
				spacer = adds + j
			}
			selected := ""
			if id == sid {
				selected = "selected"
			}
			nStr := ""
			parentIDInt, ok := value["ParentId"].(int)
			if !ok {
				parentIDInt = int(value["ParentId"].(int64))
			}
			if 0 == parentIDInt && strGroup != "" {
				nStr = strGroup
			} else {
				nStr = str
			}

			//id orm转换可能是int或者int64类型，兼容
			idInt, ok := value["Id"].(int)
			if !ok {
				idInt64, ok := value["Id"].(int64)
				if ok {
					idInt = int(idInt64)
					nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
			}

			levelInt, ok := value["Level"].(int)
			if !ok {
				levelInt64, ok := value["Level"].(int64)
				if ok {
					levelInt = int(levelInt64)
					nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
			}

			sortIDInt, ok := value["SortId"].(int)
			if !ok {
				sortIDInt64, ok := value["SortId"].(int64)
				if ok {
					sortIDInt = int(sortIDInt64)
					nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIDInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIDInt))
			}

			parentIDNodeStringValue, ok := value["ParentIdNode"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$parent_id_node", parentIDNodeStringValue)
			}

			nStr = strings.ReplaceAll(nStr, "$spacer", spacer)
			nStr = strings.ReplaceAll(nStr, "$selected", selected)

			nameValue, ok := value["Name"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$name", nameValue)
			}

			urlValue, ok := value["Url"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$url", urlValue)
			}

			nStr = strings.ReplaceAll(nStr, "$parent_id", strconv.Itoa(parentIDInt))

			iconValue, ok := value["Icon"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$icon", iconValue)
			}

			isShowValue, ok := value["IsShow"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$is_show", isShowValue)
			}

			logMethodValue, ok := value["LogMethod"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$log_method", logMethodValue)
			}

			strManageValue, ok := value["StrManage"].(string)
			if ok {
				strManageValue = strings.ReplaceAll(strManageValue, "\\", "")
				nStr = strings.ReplaceAll(nStr, "$str_manage", strManageValue)
			}

			adminTreeService.Ret += nStr
			adminTreeService.GetTree(id, str, sid, adds+k+space, strGroup)

			number++
		}

	}

	return adminTreeService.Ret
}

// Menu 菜单选择 select树形选择
func (adminTreeService *AdminTreeService) Menu(selected int, currentId int) string {
	var adminMenuService AdminMenuService
	result := adminMenuService.Menu(currentId)
	resultKey := make(map[int]orm.Params)
	if result != nil {
		for _, r := range result {
			idInt, ok := r["Id"].(int)
			if !ok {
				idInt = int(r["Id"].(int64))
			}
			resultKey[idInt] = r
			if idInt == selected {
				resultKey[idInt]["selected"] = "selected"
			} else {
				resultKey[idInt]["selected"] = ""
			}
		}

		str := `<option value='$id' $selected >$spacer $name</option>`
		adminTreeService.initTree(resultKey)
		return adminTreeService.GetTree(0, str, selected, "", "")
	}
	return ""
}

// AdminMenuTree 生成菜单树
func (adminTreeService *AdminTreeService) AdminMenuTree() string {
	var adminMenuService AdminMenuService
	adminMenus := adminMenuService.AllMenu()
	if adminMenus != nil {
		result := make(map[int]orm.Params)
		var adminTreeService AdminTreeService
		for _, adminMenu := range adminMenus {
			n := adminMenu.Id
			//初始化orm.Params map类型
			if result[n] == nil {
				result[n] = make(orm.Params)
			}

			result[n]["Id"] = adminMenu.Id
			result[n]["ParentId"] = adminMenu.ParentId
			result[n]["Name"] = adminMenu.Name
			result[n]["Url"] = adminMenu.Url
			result[n]["Icon"] = adminMenu.Icon
			result[n]["IsShow"] = adminMenu.IsShow
			result[n]["SortId"] = adminMenu.SortId

			result[n]["Level"] = adminTreeService.GetLevel(adminMenu.Id, result, 0)
			if adminMenu.ParentId > 0 {
				result[n]["ParentIdNode"] = ` class="child-of-node-` + strconv.Itoa(adminMenu.ParentId) + `"`
			} else {
				result[n]["ParentIdNode"] = ""
			}
			result[n]["StrManage"] = `<a href="/admin/admin_menu/edit?id=` + strconv.Itoa(adminMenu.Id) + `" class="btn btn-primary btn-xs" title="修改" data-toggle="tooltip"><i class="fa fa-pencil"></i></a> <a class="btn btn-danger btn-xs AjaxButton" data-id="` + strconv.Itoa(adminMenu.Id) + `" data-url="del"  data-confirm-title="删除确认" data-confirm-content=\'您确定要删除ID为 <span class="text-red"> ` + strconv.Itoa(adminMenu.Id) + ` </span> 的数据吗\'  data-toggle="tooltip" title="删除"><i class="fa fa-trash"></i></a>`
			if adminMenu.IsShow == 1 {
				result[n]["IsShow"] = "显示"
			} else {
				result[n]["IsShow"] = "隐藏"
			}
			result[n]["LogMethod"] = adminMenu.LogMethod
		}
		str := `<tr id='node-$id' data-level='$level' $parent_id_node><td><input type='checkbox' onclick='checkThis(this)'
                     name='data-checkbox' data-id='$id' class='checkbox data-list-check' value='$id' placeholder='选择/取消'>
                    </td><td>$id</td><td>$spacer$name</td><td>$url</td>
                    <td>$parent_id</td><td><i class='fa $icon'></i><span>($icon)</span></td>
                    <td>$sort_id</td><td>$is_show</td><td>$log_method</td><td class='td-do'>$str_manage</td></tr>`

		adminTreeService.initTree(result)

		return adminTreeService.GetTree(0, str, 0, "", "")

	}
	return ""
}

// AuthorizeHtml 生成授权html
func (adminTreeService *AdminTreeService) AuthorizeHtml(menu map[int]orm.Params, authMenus []string) string {
	for id := range menu {
		if utils.InArrayForString(authMenus, strconv.Itoa(id)) {
			menu[id]["Checked"] = " checked"
		} else {
			menu[id]["Checked"] = ""
		}
		levelInt := adminTreeService.GetLevel(id, menu, 0)
		menu[id]["Level"] = levelInt
		menu[id]["Width"] = 100 - levelInt
	}

	adminTreeService.initTree(menu)

	adminTreeService.Text = make(map[string]interface{})
	adminTreeService.Text["other"] = `<label class='checkbox'  >
                        <input $checked  name='url[]' value='$id' level='$level'
                        onclick='javascript:checkNode(this);' type='checkbox'>
                       $name
                   </label>`

	adminTreeService.Text["0"] = []string{
		`<dl class='checkMod'>
                    <dt class='hd'>
                        <label class='checkbox'>
                            <input $checked name='url[]' value='$id' level='$level'
                             onclick='javascript:checkNode(this);'
                             type='checkbox'>
                            $name
                        </label>
                    </dt>
                    <dd class='bd'>`,
		`</dd></dl>`,
	}

	adminTreeService.Text["1"] = []string{
		`
                        <div class='menu_parent'>
                            <label class='checkbox'>
                                <input $checked  name='url[]' value='$id' level='$level'
                                onclick='javascript:checkNode(this);' type='checkbox'>
                               $name
                            </label>
                        </div>
                        <div class='rule_check' style='width: $width%;'>`,
		`</div><span class='child_row'></span>`,
	}

	return adminTreeService.getAuthTreeAccess(0)
}

// getAuthTreeAccess 获取权限树
func (adminTreeService *AdminTreeService) getAuthTreeAccess(myID int) string {
	nStr := ""
	child := adminTreeService.getChild(myID)

	if len(child) > 0 {
		//取key最小的一个，防止for range随机取，导致每次菜单顺序不同
		sortID := 99999
		for k := range child {
			if k < sortID {
				sortID = k
			}
		}
		level := make(map[string]interface{})
		level = child[sortID]

		//child排序
		var ids []int
		for id := range child {
			ids = append(ids, id)
		}
		sort.Ints(ids)

		var text []string
		if v, ok := adminTreeService.Text[strconv.Itoa(level["Level"].(int))]; ok {
			text = v.([]string)
		} else {
			text = adminTreeService.Text["1"].([]string)
		}

		for _, id := range ids {
			k := id
			v := child[id]
			if len(adminTreeService.getChild(k)) > 0 {
				nStr = text[0]
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(v["Id"].(int)))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(v["Level"].(int)))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				nStr = strings.ReplaceAll(nStr, "$width", strconv.Itoa(v["Width"].(int)))
				adminTreeService.Html += nStr

				adminTreeService.getAuthTreeAccess(k)

				nStr = text[1]
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(v["Id"].(int)))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(v["Level"].(int)))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				nStr = strings.ReplaceAll(nStr, "$width", strconv.Itoa(v["Width"].(int)))
				adminTreeService.Html += nStr

			} else {
				nStr = adminTreeService.Text["other"].(string)
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(v["Id"].(int)))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(v["Level"].(int)))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				adminTreeService.Html += nStr
			}
		}

		return adminTreeService.Html
	}

	return ""
}
