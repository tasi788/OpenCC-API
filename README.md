---
description: 快速將中文字在簡繁之中切換。
---

# Strings

{% api-method method="post" host="https://strings.hexelf.dev" path="" %}
{% api-method-summary %}
 翻譯
{% endapi-method-summary %}

{% api-method-description %}
直接 **POST** 來取得翻譯內容。
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="target" type="string" required=true %}
 翻譯至 `hk2s, s2hk, s2t, s2tw, s2twp, t2hk, t2s, t2tw, tw2s, tw2sp`
{% endapi-method-parameter %}

{% api-method-parameter name="text" type="string" required=true %}
輸入字串 
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
執行分以下幾種
{% endapi-method-response-example-description %}

```javascript
// 正常操作結束
{
    "Status": true,
    "Text": "我終於會繁體惹。"
}

// 未給參數
{    
    "Status": false,
    "Text":"Args Not Found "
}

// 未輸入正確翻譯字典名稱
{
    "Status": false,
    "Text":"Reading Dictionary Failed!"
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



> Photo Credit:  made by [monkik](https://www.flaticon.com/authors/monkik)

