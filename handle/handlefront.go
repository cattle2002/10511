package handle

//func ConvertAlgoModelToHttpAlgoModel(models []model.AlgoModel) []protocol.HttpAlgoModel {
//	var result []protocol.HttpAlgoModel
//
//	for _, algo := range models {
//		httpAlgo := protocol.HttpAlgoModel{
//			// 进行结构体字段赋值
//			ID:       algo.ID,
//			CreateAt: algo.CreatedAt,
//			UpdateAt: algo.UpdatedAt,
//			Position: algo.Position,
//			FileName: algo.FileName,
//			AlName:   algo.AlName,
//			FunName:  algo.FunName,
//			Details:  algo.Details,
//			Type:     algo.Type,
//			Inputs:   algo.Inputs,
//			OutPuts:  algo.OutPuts,
//		}
//
//		result = append(result, httpAlgo)
//	}
//
//	return result
//}
//func ListAlgo(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		pageNum := r.PostFormValue("PageNum")
//		pageSize := r.PostFormValue("PageSize")
//		pn, e := strconv.Atoi(pageNum)
//		if e != nil {
//			fmt.Println(e)
//		}
//		ps, e1 := strconv.Atoi(pageSize)
//		if e1 != nil {
//			fmt.Println(e1)
//		}
//		fmt.Println(pn, ps)
//		am, err := model.FindAlgoPage(pn, ps)
//		if err != nil {
//			fmt.Println(e)
//		}
//		data := ConvertAlgoModelToHttpAlgoModel(am)
//
//		res := protocol.HttpListAlgoResponse{
//			Code: 200,
//			Msg:  "success",
//			Data: ni,
//		}
//		b, err2 := json.Marshal(res)
//		if err2 != nil {
//			fmt.Println(err2)
//		}
//		_, e = w.Write(b)
//		if e != nil {
//			fmt.Println("err:", err)
//		}
//	} else {
//		res := protocol.HttpListAlgoResponse{Code: http.StatusMethodNotAllowed, Msg: "method is illeagl", Data: nil}
//		marshal, err := json.Marshal(res)
//		if err != nil {
//			fmt.Println("marshal error:", err)
//		}
//		w.Write(marshal)
//	}
//}
