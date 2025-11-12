package template

import (
	"fmt"

	"github.com/carfdev/carfdev-emailsvc/internal/types"
)

func ContactRequestTemplate(req *types.SendContactRequest) string {
	message := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nueva Solicitud de Contacto</title>
</head>
<body style="margin: 0; padding: 0; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif; background-color: #f4f4f5; line-height: 1.6;">
    <table role="presentation" style="width: 100%%; border-collapse: collapse; background-color: #f4f4f5;">
        <tr>
            <td style="padding: 40px 20px;">
                <table role="presentation" style="max-width: 600px; margin: 0 auto; background-color: #ffffff; border-radius: 8px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);">
                    <!-- Header -->
                    <tr>
                        <td style="padding: 40px 40px 30px; text-align: center; background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%); border-radius: 8px 8px 0 0;">
                            <h1 style="margin: 0; color: #ffffff; font-size: 28px; font-weight: 600;">
                                Nueva Solicitud de Contacto
                            </h1>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px;">
                            <p style="margin: 0 0 24px; color: #52525b; font-size: 16px;">
                                Has recibido una nueva solicitud de contacto desde tu sitio web.
                            </p>
                            
                            <!-- Info Grid -->
                            <table role="presentation" style="width: 100%%; border-collapse: collapse;">
                                <tr>
                                    <td style="padding: 16px; background-color: #f9fafb; border-radius: 6px; margin-bottom: 16px;">
                                        <table role="presentation" style="width: 100%%;">
                                            <tr>
                                                <td style="padding-bottom: 12px;">
                                                    <strong style="color: #3f3f46; font-size: 14px; text-transform: uppercase; letter-spacing: 0.5px;">
                                                        👤 Información Personal
                                                    </strong>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td style="padding: 8px 0; border-top: 1px solid #e4e4e7;">
                                                    <span style="color: #71717a; font-size: 14px;">Nombre:</span><br>
                                                    <span style="color: #18181b; font-size: 16px; font-weight: 500;">%s %s</span>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td style="padding: 8px 0; border-top: 1px solid #e4e4e7;">
                                                    <span style="color: #71717a; font-size: 14px;">Email:</span><br>
                                                    <a href="mailto:%s" style="color: #667eea; font-size: 16px; text-decoration: none;">%s</a>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td style="padding: 8px 0; border-top: 1px solid #e4e4e7;">
                                                    <span style="color: #71717a; font-size: 14px;">Empresa:</span><br>
                                                    <span style="color: #18181b; font-size: 16px; font-weight: 500;">%s</span>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>

																<tr>
    															<td style="height: 16px;"></td>
																</tr>
                                
                                <tr>
                                    <td style="padding: 16px; background-color: #f9fafb; border-radius: 6px; margin-top: 16px;">
                                        <table role="presentation" style="width: 100%%;">
                                            <tr>
                                                <td style="padding-bottom: 12px;">
                                                    <strong style="color: #3f3f46; font-size: 14px; text-transform: uppercase; letter-spacing: 0.5px;">
                                                        💼 Detalles del Proyecto
                                                    </strong>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td style="padding: 8px 0; border-top: 1px solid #e4e4e7;">
                                                    <span style="color: #71717a; font-size: 14px;">Tipo de Proyecto:</span><br>
                                                    <span style="color: #18181b; font-size: 16px; font-weight: 500;">%s</span>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td style="padding: 8px 0; border-top: 1px solid #e4e4e7;">
                                                    <span style="color: #71717a; font-size: 14px;">Presupuesto:</span><br>
                                                    <span style="color: #18181b; font-size: 16px; font-weight: 500;">%s</span>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>

																<tr>
    															<td style="height: 16px;"></td>
																</tr>
                                
                                <tr>
                                    <td style="padding: 16px; background-color: #f9fafb; border-radius: 6px; margin-top: 16px;">
                                        <strong style="color: #3f3f46; font-size: 14px; text-transform: uppercase; letter-spacing: 0.5px; display: block; padding-bottom: 12px;">
                                            💬 Mensaje
                                        </strong>
                                        <div style="padding: 12px 0; border-top: 1px solid #e4e4e7; color: #18181b; font-size: 15px; white-space: pre-wrap;">%s</div>
                                    </td>
                                </tr>
                            </table>
                            
                            <!-- CTA Button -->
                            <table role="presentation" style="width: 100%%; margin-top: 32px;">
                                <tr>
                                    <td style="text-align: center;">
                                        <a href="mailto:%s" style="display: inline-block; padding: 14px 32px; background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%); color: #ffffff; text-decoration: none; border-radius: 6px; font-weight: 600; font-size: 16px;">
                                            Responder al Cliente
                                        </a>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="padding: 24px 40px; background-color: #fafafa; border-radius: 0 0 8px 8px; text-align: center; border-top: 1px solid #e4e4e7;">
                            <p style="margin: 0; color: #71717a; font-size: 13px;">
                                Este correo fue enviado automáticamente desde el formulario de contacto de tu sitio web.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`,
		req.FirstName, req.LastName,
		req.Email, req.Email,
		req.CompanyName,
		req.ProjectType,
		req.Budget,
		req.Message,
		req.Email,
	)

	return message
}
