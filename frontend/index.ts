import express, { Express, Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";
import dotenv from "dotenv";
import path from "path";

dotenv.config();

const app: Express = express();
const port = 3000;

app.set("view engine", "ejs");
app.use(express.static(path.join(__dirname, "public")));

app.use((req: Request, res: Response, next: NextFunction) => {
    const bearerToken = process.env.BAROSA_BEARER_AUTH;
    const secret = process.env.BAROSA_BEARER_SECRET;

    if (!bearerToken || !secret) {
        return res.status(500).send("BAROSA_BEARER_AUTH or BAROSA_BEARER_SECRET not found in environment variables");
    }

    const token = jwt.sign({ bearerToken }, secret, { expiresIn: "5h" });
    res.locals.token = token;

    next();
});

app.get("/", (req: Request, res: Response) => {
  res.render("index", { token: res.locals.token });
})

app.listen(port, () => {
  console.log(`[server]: Express is running at http://localhost:${port}`);
});
