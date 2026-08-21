# Quiadsa Screen Sharing

Compartición de pantalla simple y segura con WebRTC.

## Descripción general

**Quiadsa Screen Sharing** es una aplicación minimalista de compartición de pantalla basada en navegador. Utiliza WebRTC para la transmisión en tiempo real de igual a igual (peer-to-peer) y un único binario de Go como backend para la señalización y el servicio de archivos estáticos. No requiere plugins ni instalaciones: basta con abrir el navegador para empezar a compartir o ver una pantalla.

- **Un único binario de Go**: Sirve el frontend y gestiona toda la lógica del backend (señalización WebSocket, gestión de sesiones).
- **WebRTC**: Transmisión directa y cifrada de igual a igual entre emisor y espectadores.
- **Sin dependencias para el usuario**: Funciona en todos los navegadores modernos.

## Características

- **Compartición de pantalla instantánea**: Empieza a transmitir tu pantalla en segundos.
- **Segura**: Todas las conexiones están cifradas (WebRTC, WSS).
- **Sin cuentas ni instalaciones**: Solo tienes que compartir un código de 5 caracteres.
- **Baja latencia**: Vídeo y audio en tiempo real.
- **Multiplataforma**: Funciona en Windows, macOS, Linux y navegadores móviles (solo como espectador).

## Cómo funciona

1. El **emisor** inicia una sesión y recibe un código único de 5 caracteres.
2. Los **espectadores** introducen el código para conectarse y ver la transmisión en tiempo real.
3. Toda la señalización (oferta/respuesta/candidatos) se gestiona mediante WebSocket; después, el medio audiovisual fluye directamente de igual a igual mediante WebRTC.

---

## Ejecución en local

### Requisitos previos

- Go 1.24 o superior

### Compilar y ejecutar

```sh
git clone <url-del-repositorio>
cd screenz
go build -o quiadsa-screen-sharing
./quiadsa-screen-sharing -port=8080
```

Después, abre [http://localhost:8080](http://localhost:8080) en tu navegador.

## Licencia

Licencia MIT — ver [LICENSE](LICENSE)
