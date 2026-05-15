# art-interface — web studio for art-decoder

`art-interface` is a small Go web server that provides a **beautiful web UI** on top of the existing `art-decoder` encode/decode logic.

---

## Run

```bash
git clone https://gitea.kood.tech/hoangphuocvu/interface
cd interface
go build -o art-interface ./cmd/web
./art-interface
```

The default port is `8080`.

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

### CSS: 

Styling is served from `./cmd/web/assets/static/styles.css`.

### Use of formatting: 

Output is rendered in a `<pre>` block so whitespace, alignment, and ASCII art layout are preserved.

### Multiline input

The text box supports **multiple lines**. The server processes input **line-by-line** (like the CLI `--multi` mode).

- **Decode**: each line is decoded independently. If a line is invalid, that output line becomes the literal `Error`.
  - If **any** line fails to decode, `POST /decoder` returns **400 Bad Request** (and the output still shows a mix of decoded lines and `Error` lines).
  - If **all** lines decode successfully, it returns **202 Accepted**.
  - **Empty lines** go through the error handling, since many arts need blank lines inbetween.
- **Encode**: each line is encoded independently and `POST /decoder` returns **202 Accepted**. `Encode()` did not return any error, so if the user's input has nothing to compress, the output is **identical** with the input, also no `Error` got caught.

## Design Notes

- `POST /decoder` handles both decoding and encoding functionality. The `decoderHandler` function receives mode data from the form and decides which operation to perform.
- The result of `POST /decoder` is rendered and appended to the mainpage. In other words, there is no `GET /decoder`.
- The HTML and CSS file is baked into binary and being held in `embed.FS` object (got compiled at build time). It brings single-file deployment (`art-interface` can be run anywhere), no missing file errors, faster read and tamper-proof for those UI assets.
- The **Server** has some `Timeout` fields so be mindful if users have bad internet connection.
