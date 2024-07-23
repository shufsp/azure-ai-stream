const express = require("express");
const jwt = require("jsonwebtoken");
const dotenv = require("dotenv");
const path = require("path");

dotenv.config();

const app = express();
const port = process.env.BAROSA_CLIENT_PORT;

app.set("view engine", "ejs");
app.use(express.static(path.join(__dirname, "public")));

app.use((req, res, next) => {
    const bearerToken = process.env.BAROSA_BEARER_AUTH;
    const secret = process.env.BAROSA_BEARER_SECRET;

    if (!bearerToken || !secret) {
        return res.status(500).send("BAROSA_BEARER_AUTH or BAROSA_BEARER_SECRET not found in environment variables");
    }

    const token = jwt.sign({ bearerToken }, secret, { expiresIn: "5h" });
    res.locals.token = token;

    next();
});

app.get("/", (req, res) => {
  res.render("index", { 
    token:            res.locals.token, 
    avifQuality:      req.query["avifQuality"] || undefined,
    avifAlphaQuality: req.query["avifAlphaQuality"] || undefined,
    avifSpeed:        req.query["avifSpeed"] || undefined,
    window:           req.query["window"] || undefined,  
    method:           req.query["method"] || undefined,
    features:         req.query["features"] || undefined,
  });
});

app.listen(port, () => {
  console.log(`[server]: Express is running at http://localhost:${port}`);
});
