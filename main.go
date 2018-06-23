package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"excel/tools"
	"sync"
)

func init(){
	log.SetFlags(log.Ldate|log.Lshortfile)
}


var (
	all = make([]string,0,16)
	excelSync sync.Mutex
	wg sync.WaitGroup
	//xlsx *excelize.File
)

//var begin,endLine = 0,0
const  (
	SHEET_NAME = "Sheet1"

)
//30 w -- 1022.036330 s 5个 goroutine 16min
//10w ----- 181.467648 2个 goroutine  3 min

func main() {
	start := time.Now()
	//log.Println(len(all),all)
	//多少列
	for i:= 1;i<=16;i++{
		toChar(i)
	}
	//log.Println(len(all),all)
	//return
	sum := 100000
	per := 50000
	count := sum/per
	begin,endLine := 0,0

	for j := 1; j<= count;j++{
		if j == 1{
			begin = 0
			endLine = per*j
		}else{
			begin = per*(j-1)
			endLine = per*j
		}

		if  (j == count) &&(endLine < sum) && (endLine * count > sum){
			endLine = sum
		}
		log.Println(begin,endLine)
		wg.Add(1)
		go writeOne("filename"+strconv.Itoa(begin)+"-"+strconv.Itoa(endLine)+".xlsx",begin,endLine)
	}

	wg.Wait()
	end := time.Now()
	log.Println(tools.Sub(start,end))
}



func test()  {

	start := time.Now()
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet(SHEET_NAME)
	//多少列
	for i:= 1;i<=16;i++{
		toChar(i)
	}

	rand.Seed(time.Now().Unix())
	header := []string{
		"id",
		"供应商名称",
		"URL",
		"数目",
		"归属",
		"投票",
		"分类",
		"其他",
		"属性",
		"supplier_id",
		"supplier_cate",
		"create_time",
		"is_del",
		"update_time",
		"status",
		"category",
	}
	//抬头
	for k,v := range all{
		//转换为 双引号的字符串
		xlsx.SetCellValue(SHEET_NAME,v+"1",header[k])
	}

	sum := 300000
	per := 10000
	count := sum/per
	begin,endLine1 := 0,0

	//内容 2 表示从第二行开始执行
	for j := 2; j<= count;j++{
		if j == 2{
			begin = 2
			endLine1 = per*j
		}else{
			begin = per*(j-1)
			endLine1 = per*j
		}
		wg.Add(1)
		//log.Println(begin,endLine1)
		go write(xlsx,begin,endLine1)
	}

	wg.Wait()
	xlsx.SetActiveSheet(index)
	err := xlsx.SaveAs("./goods-44.xlsx")
	if err !=nil {
		log.Println(err)
	}
	end := time.Now()
	log.Println(tools.Sub(start,end))
}

func write(xlsx *excelize.File,startLine int,endLine int) {
	//内容
	for i := startLine;i< endLine;i++{
		for _,v := range all{
			intRandom := rand.Intn(100)
			value := "f测试咨询处 "+strconv.Itoa(intRandom)
			//excelSync.Lock()
			xlsx.SetCellValue(SHEET_NAME,v+strconv.Itoa(i),value)
			//excelSync.Unlock()
		}
	}
}

func toChar(i int)  {
	//tmp := fmt.Sprintf("%s",rune('A' - 1 + i))
	tmp := string(rune('A' - 1 + i))
	all = append(all,tmp)
}



func writeOne(fileName string,begin,endLine int)  {

	defer wg.Done()
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(SHEET_NAME)
	//多少列
	//all := all[:0:16]
	//老出错！
	//for i:= 1;i<=16;i++{
	//	toChar(i)
	//}
	rand.Seed(time.Now().Unix())
	header := []string{
		"id",
		"供应商名称",
		"URL",
		"数目",
		"归属",
		"投票",
		"分类",
		"其他",
		"属性",
		"supplier_id",
		"supplier_cate",
		"create_time",
		"is_del",
		"update_time",
		"status",
		"category",
	}

	//log.Println(len(all),header)
	//return
	//抬头
	for k,v := range all{
		//转换为 双引号的字符串
		pos := fmt.Sprintf("%s%s",v,strconv.Itoa(1))
		//fmt.Printf("%v,%s\n",pos,header[k])
		xlsx.SetCellValue(SHEET_NAME,pos,header[k])
	}

	end := endLine - begin
	for i := 2;i< end;i++{
		for _,v := range all{
			intRandom := rand.Intn(100)
			value := "f测试咨询处 "+strconv.Itoa(intRandom)
			//excelSync.Lock()
			pos := fmt.Sprintf("%s%s",v,strconv.Itoa(i))
			xlsx.SetCellValue(SHEET_NAME,pos,value)
			//excelSync.Unlock()
		}
	}

	//write(xlsx,begin,endLine)
	xlsx.SetActiveSheet(index)
	err := xlsx.SaveAs(fileName)
	if err !=nil {
		log.Println(err)
	}
}