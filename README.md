# art-interface — web studio for art-decoder

`art-interface` is a small Go web server that provides a **beautiful web UI** on top of the existing `art-decoder` encode/decode logic.

---

## Run

```bash
git clone ...
cd art-interface
go run ./cmd/web --addr :8080
```

Open `http://localhost:8080`.

---

## Endpoints

- `GET /` → **200 OK** (main page)
- `POST /decoder` → **202 Accepted** on success, **400 Bad Request** on malformed input

Any other request returns an appropriate status code (e.g. `404 Not Found`, `405 Method Not Allowed`).

---

## Extras

### Encode mode: 

Choose **Encode** in the UI radio button. The server uses `art-decoder/utils.Encode` (nothing is computed in the browser). It only encodes when doing that actually *reduce* the input length.

### PaintLine (colored output): 

Colorized the output when *Decoding* using RGB in HTML format. Off by default. Enable with `--paint`:

```bash
cd art-interface
go run ./cmd/web --addr :8080 --paint
```

### CSS: 

Styling is served from `./art-interface/cmd/web/assets/static/styles.css`.

### Use of formatting: 

Output is rendered in a `<pre>` block so whitespace, alignment, and ASCII art layout are preserved.

### Multiline input

The text box supports **multiple lines**. The server processes input **line-by-line** (like the CLI `--multi` mode).

- **Decode**: each line is decoded independently. If a line is invalid, that output line becomes the literal `Error`.
  - If **any** line fails to decode, `POST /decoder` returns **400 Bad Request** (and the output still shows a mix of decoded lines and `Error` lines).
  - If **all** lines decode successfully, it returns **202 Accepted**.
- **Encode**: each line is encoded independently and `POST /decoder` returns **202 Accepted**.

