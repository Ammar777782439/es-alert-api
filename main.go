package main

import (
    "fmt"
    "github.com/elastic/go-elasticsearch/v8"
)

func main() {
    // إنشاء عميل Elasticsearch
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        fmt.Printf("خطأ في إنشاء العميل: %s\n", err)
        return
    }

    // اختبار الاتصال عن طريق الحصول على معلومات الكلاستر
    res, err := es.Info()
    if err != nil {
        fmt.Printf("خطأ في Info: %s\n", err)
        return
    }
    defer res.Body.Close()

    fmt.Println(res)
}
