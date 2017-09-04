package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/howie6879/NIYT/fetcher"
	"github.com/modood/table"
)

type novleDemo struct {
	Index int
	Name  string
	URL   string
}

type chapterDemo struct {
	Index int
	Title string
	Href  string
}

// Token from https://github.com/amyhaber/cnki-downloader/blob/master/main.go getInputString
func getInputString() string {
	buf := bufio.NewReader(os.Stdin)
	s, err := buf.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(s)
}

func main() {
	color.Cyan("******************************************************************************\n")
	color.Cyan("****                Read the novel in your terminal - NIYT                ****\n")
	color.Cyan("******************************************************************************\n")
	fmt.Println()
	fmt.Println("**** 请先输入小说名")
	fmt.Println("**** show : 显示此时资源 如书源以及最新章节")
	fmt.Println("**** get  : 如 get 0 ，获取此资源并显示")
	fmt.Println("****")
	defer func() {
		color.Yellow("下次再见^_^\n")
	}()
	for {
		fmt.Fprintf(color.Output, "$ %s", color.CyanString("请输入小说名~ "))
		name := getInputString()
		if len(name) == 0 {
			continue
		}
		query := name + " 小说 最新章节"
		resultData, _ := fetcher.FetchResult(query)
		if len(resultData) > 0 {
			var novelData []novleDemo
			for index, data := range resultData {
				novelData = append(novelData, novleDemo{Index: index, Name: data.Title, URL: data.URL})
			}
			table.Output(novelData)
			for {
				flag := false
				fmt.Fprintf(color.Output, "$ %s", color.CyanString(name+" 源 ~ "))
				command := getInputString()
				cmdSplit := strings.Split(command, " ")
				switch strings.ToLower(cmdSplit[0]) {
				case "show":
					{
						table.Output(novelData)
					}
				case "get":
					{
						id, isOk := idJudge(cmdSplit)
						if !isOk {
							break
						}
						if id >= uint64(len(resultData)) {
							color.Red("请输入正确编号")
							break
						}
						resultData[id].FetchChapters()
						if len(resultData[id].Chapters) > 0 {
							var chapterData []chapterDemo
							for index, data := range resultData[id].Chapters {
								chapterData = append(chapterData, chapterDemo{Index: index, Title: data.ChapterName, Href: data.Href})
							}
							table.Output(chapterData)
							for {
								chapterFlag := false
								fmt.Fprintf(color.Output, "$ %s", color.CyanString(name+" 章节 ~ "))
								chapterCommand := getInputString()
								chapterCmdSplit := strings.Split(chapterCommand, " ")
								switch strings.ToLower(chapterCmdSplit[0]) {
								case "show":
									{
										table.Output(chapterData)
									}
								case "get":
									{
										chapterID, isOk := idJudge(chapterCmdSplit)
										if !isOk {
											break
										}
										if chapterID >= uint64(len(resultData[id].Chapters)) {
											color.Red("请输入正确编号")
											break
										}
										resultData[id].Chapters[chapterID].FetchContent()
										fmt.Println(resultData[id].Chapters[chapterID].Content)
									}
								case "q":
									{
										chapterFlag = true
									}
								}
								if chapterFlag {
									break
								}
							}
						}
					}
				case "q":
					{
						flag = true
					}
				}
				if flag {
					break
				}
			}
		} else {
			fmt.Println("暂无结果，请重试！")
		}
	}
}

func idJudge(cmdSplit []string) (uint64, bool) {
	if len(cmdSplit) < 2 {
		color.Red("输入格式不对")
		return 0, false
	}
	id, err := strconv.ParseUint(cmdSplit[1], 10, 32)
	if err != nil {
		fmt.Fprintf(color.Output, "输入格式不对 %s\n", color.RedString(err.Error()))
		return 0, false
	}
	return id, true
}
