package cmd

//import (
//	"corrector/app"
//
//	"github.com/spf13/cobra"
//)
//
//func start(app *app.Server) *cobra.Command {
//	return &cobra.Command{
//		Use:   "start",
//		Short: "Start showcase web-server",
//		Run: func(cmd *cobra.Command, args []string) {
//			app.Init()        // Инициируем приложение и устанавливаем соединение с БД
//			defer app.Close() // Откладываем закрытие соединения
//
//			app.Start() // Запускаем сервер
//		},
//	}
//}
