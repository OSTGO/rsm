package operate

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type RssCfgMap map[string]string

func fileCheck(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println(filePath, " not exit ,create fail!")
		}
		fmt.Println(filePath, " not exit ,successful create ", filePath)
	}
}

func (rssMap *RssCfgMap) readCfgFileAsMap(filePath string) {
	fileCheck(filePath)
	config := viper.New()
	config.SetConfigFile(filePath)
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = config.UnmarshalKey("rssList", &rssMap)
	if err != nil {
		panic(fmt.Errorf("UnmarshKey error : %s \n", err))
	}
	return
}

func (rssMap RssCfgMap) deleteRecodeByKey(name []string) {
	for _, v := range name {
		delete(rssMap, v)
	}
}

func (rssMap *RssCfgMap) deleteRecodeByValue(value []string) {
	exchangedRssMapIn, _ := exchangeMapKV(*rssMap)
	exchangedRssMapIn.deleteRecodeByKey(value)
	*rssMap, _ = exchangeMapKV(exchangedRssMapIn)
}

func (rssMap *RssCfgMap) mergeRecord(recordsList ...map[string]string) {
	// add records or merge records can use this func
	if len(*rssMap) == 0 {
		*rssMap = RssCfgMap{}
	}
	for _, records := range recordsList {
		for name, url := range records {
			(*rssMap)[name] = url
		}
	}
	exchangedRssMap, flag := exchangeMapKV(*rssMap)
	if flag {
		*rssMap, _ = exchangeMapKV(exchangedRssMap)
		return
	}
	return
}

func FileAddRecords(records map[string]string, files ...string) {
	for _, file := range files {
		temRssMap := RssCfgMap{}
		temRssMap.readCfgFileAsMap(file)
		temRssMap.mergeRecord(records)
		temRssMap.WriteCfgFileList(file)
	}
}

func FileDelRecords(keys, values []string, files ...string) {
	for _, file := range files {
		temRssMap := RssCfgMap{}
		temRssMap.readCfgFileAsMap(file)
		temRssMap.deleteRecodeByKey(keys)
		temRssMap.deleteRecodeByValue(values)
		temRssMap.WriteCfgFileList(file)
	}
}

func (rssMap *RssCfgMap) MergeCfgFile(cfgFiles ...string) {
	for _, cfgFile := range cfgFiles {
		var rssMapTem = RssCfgMap{}
		rssMapTem.readCfgFileAsMap(cfgFile)
		rssMap.mergeRecord(rssMapTem)
	}
}

func (rssMap RssCfgMap) writeCfgFile(filePath string) {
	v := viper.New()
	v.Set("rssList", rssMap)
	err := v.WriteConfigAs(filePath)
	if err != nil {
		fmt.Println("Write ", filePath, "error, because: ", err)
	}
}

func (rssMap RssCfgMap) WriteCfgFileList(filePaths ...string) {
	for _, filePath := range filePaths {
		rssMap.writeCfgFile(filePath)
	}
}

func (rssMap RssCfgMap) PrintKeyValue(keyLen, valueLen int) {
	fmt.Println(" " + strings.Repeat("-", keyLen+valueLen+4))
	printKV("title", "link", keyLen, valueLen)
	fmt.Println(" " + strings.Repeat("-", keyLen+valueLen+4))
	for key, value := range rssMap {
		printKV(key, value, keyLen, valueLen)
	}
	fmt.Println(" " + strings.Repeat("-", keyLen+valueLen+4))
}

func printKV(key, value string, keyNum, valueNum int) {
	fmt.Println(beautifulString(key, keyNum) + beautifulString(value, valueNum))
}

func beautifulString(str string, outLen int) string {
	offset := outLen - len(str)
	var front int
	var back int
	if offset >= 1 {
		front = offset / 2
		back = offset/2 + offset%2
	}
	return fmt.Sprintf("%s%s%s%s%s", "|", strings.Repeat(" ", front), str, strings.Repeat(" ", back), " |")
}

func exchangeMapKV(mapIn RssCfgMap) (mapOut RssCfgMap, sizeChangeFlag bool) {
	//if mapIn have duplicate value, sizeChangeFlag will be true
	mapOut = map[string]string{}
	for k, v := range mapIn {
		mapOut[v] = k
	}
	if len(mapIn) != len(mapOut) {
		sizeChangeFlag = true
	}
	return
}
