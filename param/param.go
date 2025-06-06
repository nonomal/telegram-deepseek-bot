package param

import (
	"github.com/cohesion-org/deepseek-go"
	"github.com/sashabaranov/go-openai"
)

const (
	DeepSeek      = "deepseek"
	DeepSeekLlava = "deepseek-ollama"

	Gemini                        = "gemini"
	ModelGemini25Pro       string = "gemini-2.5-pro"
	ModelGemini25Flash     string = "gemini-2.5-flash"
	ModelGemini20Flash     string = "gemini-2.0-flash"
	ModelGemini20FlashLite string = "gemini-2.0-flash-lite"
	ModelGemini15Pro       string = "gemini-1.5-pro"
	ModelGemini15Flash     string = "gemini-1.5-flash"
	ModelGemini10Ultra     string = "gemini-1.0-ultra"
	ModelGemini10Pro       string = "gemini-1.0-pro"
	ModelGemini10Nano      string = "gemini-1.0-nano"

	// 特定功能模型
	ModelGeminiFlashPreviewTTS string = "gemini-flash-preview-tts"
	ModelGeminiEmbedding       string = "gemini-embedding"
	ModelImagen3               string = "imagen-3"
	ModelVeo2                  string = "veo-2"

	OpenAi = "openai"

	LLAVA = "llava:latest"

	ImageTokenUsage = 10000
	VideoTokenUsage = 20000
)

var (
	GeminiModels = map[string]bool{
		ModelGemini25Pro:       true,
		ModelGemini25Flash:     true,
		ModelGemini20Flash:     true,
		ModelGemini20FlashLite: true,
		ModelGemini15Pro:       true,
		ModelGemini15Flash:     true,
		ModelGemini10Ultra:     true,
		ModelGemini10Pro:       true,
		ModelGemini10Nano:      true,
	}

	DeepseekModels = map[string]bool{
		deepseek.DeepSeekChat:     true,
		deepseek.DeepSeekReasoner: true,
		deepseek.DeepSeekCoder:    true,
	}

	OpenAIModels = map[string]bool{
		openai.GPT3Dot5Turbo0125:       true,
		openai.O1Mini:                  true,
		openai.O1Mini20240912:          true,
		openai.O1Preview:               true,
		openai.O1Preview20240912:       true,
		openai.O1:                      true,
		openai.O120241217:              true,
		openai.O3Mini:                  true,
		openai.O3Mini20250131:          true,
		openai.GPT432K0613:             true,
		openai.GPT432K0314:             true,
		openai.GPT432K:                 true,
		openai.GPT40613:                true,
		openai.GPT40314:                true,
		openai.GPT4o:                   true,
		openai.GPT4o20240513:           true,
		openai.GPT4o20240806:           true,
		openai.GPT4o20241120:           true,
		openai.GPT4oLatest:             true,
		openai.GPT4oMini:               true,
		openai.GPT4oMini20240718:       true,
		openai.GPT4Turbo:               true,
		openai.GPT4Turbo20240409:       true,
		openai.GPT4Turbo0125:           true,
		openai.GPT4Turbo1106:           true,
		openai.GPT4TurboPreview:        true,
		openai.GPT4VisionPreview:       true,
		openai.GPT4:                    true,
		openai.GPT4Dot5Preview:         true,
		openai.GPT4Dot5Preview20250227: true,
		openai.GPT3Dot5Turbo1106:       true,
		openai.GPT3Dot5Turbo0613:       true,
		openai.GPT3Dot5Turbo0301:       true,
		openai.GPT3Dot5Turbo16K:        true,
		openai.GPT3Dot5Turbo16K0613:    true,
		openai.GPT3Dot5Turbo:           true,
		openai.GPT3Dot5TurboInstruct:   true,
	}

	DeepseekLocalModels = map[string]bool{
		LLAVA:                         true,
		deepseek.AzureDeepSeekR1:      true,
		deepseek.OpenRouterDeepSeekR1: true,
		deepseek.OpenRouterDeepSeekR1DistillLlama70B: true,
		deepseek.OpenRouterDeepSeekR1DistillLlama8B:  true,
		deepseek.OpenRouterDeepSeekR1DistillQwen14B:  true,
		deepseek.OpenRouterDeepSeekR1DistillQwen1_5B: true,
		deepseek.OpenRouterDeepSeekR1DistillQwen32B:  true,
	}
)

type MsgInfo struct {
	MsgId   int
	Content string
	SendLen int
}

type ImgResponse struct {
	Code    int              `json:"code"`
	Data    *ImgResponseData `json:"data"`
	Message string           `json:"message"`
	Status  string           `json:"status"`
}

type ImgResponseData struct {
	AlgorithmBaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	ImageUrls        []string `json:"image_urls"`
	PeResult         string   `json:"pe_result"`
	PredictTagResult string   `json:"predict_tag_result"`
	RephraserResult  string   `json:"rephraser_result"`
}
