
# Gemini API Console

A simple Go console app for chatting with **Gemini 1.5 Flash**.

![Gemini 1.5 Flash console demo](./images/gemini-api-console.jpeg)

---

## Setup

1. Clone the repository:

```bash
git clone https://github.com/LeonibelDev/gemini-api-console
cd gemini-api-console
```

2. Create a `config` folder in the root directory.

3. Inside `config`, create a `.env` file containing your Gemini API key:

```env
GEMINI_API_KEY=your-api-key-here
```

Alternatively, you can set the `GEMINI_API_KEY` as a system environment variable.

---

## Run

After setting up the environment variable, build and run the app:

```bash
go build -o gemini-console
./gemini-console
```

> ⚡ Note: Make sure you have Go installed on your machine (version 1.20+ recommended).


## License

MIT
