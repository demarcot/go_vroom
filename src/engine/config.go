package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
);

type ConfigService struct {
  configMap map[string]string;
}

func (c *ConfigService) LoadConfig (path string) {
  c.configMap = make(map[string]string)

  file, err := os.Open(path);
  if (err != nil) {
    fmt.Println("Failed loading config...");
    return
  }
  defer file.Close();

  scanner := bufio.NewScanner(file);
  scanner.Split(bufio.ScanLines);

  for scanner.Scan() {
    line := scanner.Text()
    splitLine := strings.Split(line, "=");

    if len(splitLine) == 2 {
      c.configMap[splitLine[0]] = splitLine[1];
    }
  }
}

func (c *ConfigService) GetStrVal(key string, defaultVal ...string) string {
  val, ok := c.configMap[key]
  if (ok) {
    return val
  }

  return defaultVal[0];
}

func (c *ConfigService) GetIntVal(key string, defaultVal ...int) int {
  val, ok := c.configMap[key]
  if (ok) {
    intVal, err := strconv.Atoi(val)

    if (err != nil) {
      fmt.Printf("Failed str->int conversion: %s", val)
      return defaultVal[0];
    }

    return intVal
  }

  return defaultVal[0];
}

func (c *ConfigService) GetFloatVal(key string, defaultVal ...float32) float32 {
  val, ok := c.configMap[key]
  if (ok) {
    floatVal, err := strconv.ParseFloat(val, 32)

    if (err != nil) {
      fmt.Printf("Failed str->float conversion: %s", val);
      return defaultVal[0];
    }

    return float32(floatVal)
  }

  return defaultVal[0];
}
