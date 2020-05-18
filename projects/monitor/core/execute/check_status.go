package execute

import "goAction/projects/monitor/core/monitor"

// 检查Http请求的Code
// 根据监控Step的: CodeMin, CodeMinExpr, CodeMax, CodeMaxExpr
func CheckResponseStatusCode(step *monitor.Step, code int) bool {
	// 第1步：根据CodeMinExpr处理
	var result bool
	if step.CodeMinExpr == ">=" {
		if code >= step.CodeMin {
			// 只有这个情况才判断大的值
			switch step.CodeMaxExpr {
			case "<=":
				result = code <= step.CodeMax
			case "<":
				result = code < step.CodeMax
			case ">":
				result = code > step.CodeMax
			case ">=":
				result = code >= step.CodeMax
			default:
				result = true
			}
		} else {
			result = false
		}

	} else if step.CodeMinExpr == "=" {
		return code == step.CodeMin

	} else if step.CodeMinExpr == "<=" {
		return code <= step.CodeMin

	} else {
		return code < step.CodeMin
	}
	return result
}
