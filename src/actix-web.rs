use actix_web::{get, put, post, web, App, HttpServer, Responder};

#[get("/")]
async fn index() -> impl Responder {
    "Hello World!"
}

#[put("/")]
async fn put_handler() -> impl Responder {
    "PUT request received"
}

#[post("/backend-api/conversation")]
async fn post_handler(req_body: web::Json<serde_json::Value>) -> impl Responder {
    let message = &req_body.message;
    println!("{}", message);
    web::Json(json!({"msg": "2+2=4"}))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(index)
            .service(put_handler)
            .service(post_handler)
    })
    .bind("127.0.0.1:3011")?
    .run()
    .await
}
