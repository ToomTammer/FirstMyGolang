package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

// ENV หรือชื่อเต็มคือ environment variables คือค่าที่ computer แต่ละเครื่องที่เก็บคู่กับสภาพแวดล้อมของตัวเองเอาไว้ ซึ่งจะสามารถใช้งานร่วมกับ process ของ computer ได้ (มองง่ายๆเป็นตัวแปรระดับ process ของ computer)
// เช่น
// TEMP = เตัวแปรสำหรับเก็บ location ของ temporary files
// HOME = เก็บที่อยู่บ้านของแต่ละเครื่อง

func getEnv(key, fallback string) string {
  if value, exists := os.LookupEnv(key); exists {
    return value
  }
  return fallback
}

func getConfig(c *fiber.Ctx) error {
  // Example: Return a configuration value from environment variable
  secretKey := getEnv("SECRET_KEY", "defaultSecret")

  return c.JSON(fiber.Map{
    "secret_key": secretKey,
  })
}
