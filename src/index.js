const express  = require("express");
// For Getting body to parse using middle ware
//app.use(express.json()) ;
const bodyParser = require("body-parser");
const fs = require("fs");
fs.readFile("a.txt", "utf-8", function(err,data){
    
})

const app = express();
const port = 3011 

//middleware for body
app.use(bodyParser.json());

app.post("/backend-api/conversation", (req,res) => {
    const message = req.body.message;
    console.log(message);

    res.json({
        msg: "2+2=4"
    })
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})