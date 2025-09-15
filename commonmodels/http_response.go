package commonmodels

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](data T) Response[T] {
	return Response[T]{Status: "success", Data: data}
}

func Error(message string) Response[any] {
	return Response[any]{Status: "error", Message: message}
}
