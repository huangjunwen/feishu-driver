package webhook

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertJSONEqual(assert *assert.Assertions, expected, actual []byte) {

	expectedVal := map[string]interface{}{}
	actualVal := map[string]interface{}{}
	err := json.Unmarshal(expected, &expectedVal)
	assert.NoError(err)
	err = json.Unmarshal(actual, &actualVal)
	assert.NoError(err)
	assert.Equal(expectedVal, actualVal)

}

func TestDecrypter(t *testing.T) {
	assert := assert.New(t)

	// 见例子： https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
	d := newDecrypter("kudryavka")
	{
		expected := []byte(`
			{
				"uuid": "5226cd85b4d843dccee2e279d93f3ed3",
				"event": {
					"app_id": "cli_9e28cb7ba56a100e",
					"before_status": {
						"is_active": true,
						"is_frozen": true,
						"is_resigned": false
					},
					"change_time": "2020-05-20 18:33:25",
					"current_status": {
						"is_active": true,
						"is_frozen": false,
						"is_resigned": false
					},
					"employee_id": "75ge6c49",
					"open_id": "ou_2ef04637d933f798dcb92c99e845ed09",
					"tenant_key": "2d520d3b434f175e",
					"type": "user_status_change"
				},
				"token": "GzhQEyfUcx7eEungQFWtXgCbxSpUOJIb",
				"ts": "1589970805.376395",
				"type": "event_callback"
			}
		`)
		plainText, err := d.Decrypt("FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0nKabA+u/5vrUD76flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL45XryxHtiZ7JdAlPmjVsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0k3M+AvzMcw=")
		assert.NoError(err)
		assertJSONEqual(assert, expected, plainText)

		fmt.Printf("%q\n", plainText)
	}

	// 失败例子
	for i, testCase := range []struct {
		B64CipherText string
	}{
		// base 64 decode error
		{B64CipherText: "F"},
		// invalid iv (只有 8 个字节数据)
		{B64CipherText: "FIAfJPGRmFY="},
		// 无数据
		{B64CipherText: "FIAfJPGRmFZWkaxPQ1XrJQ=="},
		// 数据只有 8 个字节，必须是 16 （blockSize）的倍数
		{B64CipherText: "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfL"},
	} {
		_, err := d.Decrypt(testCase.B64CipherText)
		assert.Error(err, "test case %d", i)

		fmt.Println(err)
	}

}
