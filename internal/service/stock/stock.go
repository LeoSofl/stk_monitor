package stock

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type StockService struct {
	client *http.Client
}

func NewStockService() *StockService {
	return &StockService{
		client: &http.Client{},
	}
}

func (s *StockService) GetStockPrice(code string) (float64, error) {
	url := fmt.Sprintf("http://hq.sinajs.cn/list=%s", code)

	// 创建新的请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("创建请求失败: %v", err)
	}

	// 添加请求头
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Add("Referer", "http://finance.sina.com.cn")

	// 发送请求
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("获取股票数据失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("读取响应失败: %v", err)
	}

	// GBK转UTF-8
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	utf8Body, err := ioutil.ReadAll(reader)
	if err != nil {
		return 0, fmt.Errorf("编码转换失败: %v", err)
	}

	// 解析数据
	data := string(utf8Body)
	parts := strings.Split(data, ",")
	if len(parts) <= 3 {
		return 0, fmt.Errorf("数据格式错误")
	}

	// 获取当前价格
	price, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		return 0, fmt.Errorf("价格转换失败: %v", err)
	}

	return price, nil
}
