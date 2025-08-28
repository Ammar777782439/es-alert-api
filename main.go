package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
	 "github.com/spf13/viper"
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



	app := fiber.New()

	// 2. تعريف مسار أساسي (GET /) للرد على الطلبات
	// هذا هو الحد الأدنى المطلوب لكي يفعل الخادم شيئًا ما
	app.Get("/", func(c *fiber.Ctx) error {
		// إرسال رد نصي بسيط
		return c.SendString("مرحباً، خادم Fiber يعمل بنجاح!")
	})

	// 3. تشغيل الخادم على المنفذ 8080
	log.Println("خادم Fiber يستمع على المنفذ 8080")
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("فشل تشغيل الخادم: %s", err)
	}
}
