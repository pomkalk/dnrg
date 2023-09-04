use actix_web::{App, get, HttpResponse, HttpServer, Responder, web};
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct Query {
    name: Option<String>
}

#[derive(Deserialize, Serialize)]
struct Resp {
    message: String
}


#[get("/api2/greeting")]
async fn greeting (query: web::Query<Query>) -> impl Responder {
    let name = match query.name.clone() {
        Some(x) => x,
        None => String::default()
    };
    let msg = format!("Hello, {} from rust", name);
    let r = Resp{message: msg};
    HttpResponse::Ok().json(r)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new().service(greeting)
    }).bind(("127.0.0.1", 4001))?.run().await
}
