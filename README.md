# Vem

Простая утилита для извлечения аудио из видео или обрезки видео файлов (`.mov`).

## Быстрый старт

**Требования:** Go и FFmpeg.

1.  **Установка:**
    ```bash
    git clone https://github.com/Villhard/vem.git
    cd vem
    make install
    ```

2.  **Использование:**
    ```bash
    # Извлечение аудио
    vem -i input.mov -vn

    # Извлечение аудио с обрезкой
    vem -i input.mov -vn -ss 00:01:10 -to 00:02:45

    # Обрезка видео
    vem -i input.mov -ss 00:01:10 -to 00:02:45
    ```
    
    **Флаги:**
    - `-vn` - извлечь только аудио
    - `-ss` - время начала (HH:MM:SS)
    - `-to` - время окончания (HH:MM:SS)

## Обновление

Для обновления до последней версии:

```bash
make update
```

## Конфигурация

Настройки качества аудио можно изменить в файле `~/.config/vem/config.yaml`:

```yaml
audio:
  # Битрейт: 32k, 64k, 96k, 128k, 192k, 256k, 320k
  bitrate: "64k"
  
  # Частота дискретизации: 8000, 16000, 22050, 44100, 48000
  sample_rate: 16000
```

## Удаление

```bash
make uninstall
```
