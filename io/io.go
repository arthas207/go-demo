package io

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Example() {
	var content, _ = json.Marshal("jira:\n\nhttps://bitdesk.atlassian.net/browse/CORE-187 \n\nPHANTASMA主网链 新币对接\n1、主网名称\nPhantasma主网\n\n2，网站\nhttps://www.phantasma.io/\n\n3，技术文档（包括开发文档、API文档、节点部署）。\nhttps://phantasma.gitbook.io/phantasmachain/setup-configuration/non-validator-node\n\n4，主网探索器\nngexplorer.phantasma.io\n\n5，公共节点\nbp1.phantasma.io:7077; bp1.phantasma.io:7078; bp2.phantasma.io:7077; bp2.phantasma.io:7078。\n\n6、节点快照地址\n暂未提供\n\n7，桥接\n只有主网\n8，对接代码和节点的GitHub\nhttps://github.com/phantasma-io/phantasma-ng\n\n9, 上市的最后期限\n1月5日\n10、其他支持你的主网的CEX的时间安排\n很可能是1月中旬。\n\n11、在技术上与你的主网最相似的网络是什么？\n没有\n\n12、在你们的交易中，技术团队有什么需要注意的吗（费用、税收、延迟等）？\n由于我们是一个双代币链，我们使用我们的燃料代币KCAL来处理tx，tx价格最低是10000，最小气体限制是2100\n区块延迟将是2 - 5s\n\n13、其他说明\n我们的共识机制是Tendermint。\n\nhttp://phantasma.gitbook.io\n (https://phantasma.gitbook.io/phantasmachain/setup-configuration/non-validator-node)非验证器节点\n\n 币种特殊点：\n\n1.有多个token，主币采用的是KCAL，SOUL和其他的token需要通过转token的方式进行转账。\n2.地址采用前缀+base58编码（地址类型+公钥）的方式生成，在发送交易时，不发送公钥，通过直接解码地址来获取公钥验签。\n\n3.nonce内置，不用传。\n\n4.费用默认gasPrice:100000 gasLimit:500\n\n5.tokenAddress是token的symbol，比如SOUL，而不是合约地址，因为构建交易的时候需要传入symbol。\n\n6.gasPrice/gasLimit，主网gasPrice是100000，gasLimit多数都是2100，但是实际使用不到500，暂时设置成500.\n\n7.交易很多异常上链后才会失败，状态会变成Fault，比如验签失败。\n\nsdk：https://github.com/phantasma-io/phantasma-ts\n测试网浏览器：https://exptestnet1337.vercel.app\n主网浏览器：https://ngexplorer.phantasma.io/")
	var jsonBody = []byte(`{"type":"page","title":"TEST","space":{"key":"BITDESK"},"ancestors":[{"id":4685841}],"title":"My Test Page","body":{"storage":{"value":` + string(content) + `,"representation":"storage"}}}`)
	var bodyReader = bytes.NewReader(jsonBody)
	fmt.Println("body:", string(jsonBody))
	var req, err = http.NewRequest(http.MethodPost, "https://bitdesk.atlassian.net/wiki/rest/api/content/", bodyReader)
	if err != nil {
		fmt.Errorf("post error,error:%v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	var authStr = "84426826@qq.com:nbHBYjRbCeVTAmh4u95N37DD"
	var encodeStr = base64.StdEncoding.EncodeToString([]byte(authStr))
	req.Header.Set("Authorization", "Basic "+encodeStr)
	var client = http.Client{
		Timeout: 30 * time.Second,
	}
	var resp, error = client.Do(req)
	if error != nil {
		fmt.Errorf("client error,error:%v", err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	getReq, err := http.NewRequest(http.MethodGet, "https://bitdesk.atlassian.net/wiki/rest/api/content/", bodyReader)
	if err != nil {
		fmt.Errorf("post error,error:%v", err)
	}

}
