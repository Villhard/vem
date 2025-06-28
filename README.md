# Vem

Простая утилита для извлечения аудио (`.m4a`) из видео (`.mov`).

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
    # Простая конвертация
    vem -i input.mov

    # Обрезка по времени
    vem -ss 00:01:10 -to 00:02:45 -i input.mov
    ```
    Аудиофайл `.m4a` появится в той же папке.

## Удаление

```bash
make uninstall
```
