use actix_web::{web, App, HttpServer, Responder, HttpResponse};
use serde::{Deserialize, Serialize};
use std::sync::{Mutex};

#[derive(Serialize, Deserialize, Clone)]
struct User {
    username: String,
    email: String,
}

struct AppState {
    users: Mutex<Vec<User>>,
}

async fn ping() -> impl Responder {
    HttpResponse::Ok().json("Pong")
}

async fn get_users(data: web::Data<AppState>) -> impl Responder {
    let users = data.users.lock().unwrap();
    HttpResponse::Ok().json(&*users)
}

async fn post_user(data: web::Data<AppState>, new_user: web::Json<User>) -> impl Responder {
    let mut users = data.users.lock().unwrap();
    users.push(new_user.into_inner());
    HttpResponse::Ok().json("Success")
}

async fn get_user(username: web::Path<String>, data: web::Data<AppState>) -> impl Responder {
    let users = data.users.lock().unwrap();
    if let Some(user) = users.iter().find(|u| &u.username == username.as_str()) {
        HttpResponse::Ok().json(user)
    } else {
        HttpResponse::NotFound().finish()
    }
}

async fn delete_user(username: web::Path<String>, data: web::Data<AppState>) -> impl Responder {
    let mut users = data.users.lock().unwrap();
    if let Some(pos) = users.iter().position(|u| u.username == username.as_str()) {
        users.remove(pos);
    } 
    HttpResponse::Ok().finish()
}

async fn patch_user(username: web::Path<String>, new_user: web::Json<User>, data: web::Data<AppState>) -> impl Responder {
    let mut users = data.users.lock().unwrap();
    if let Some(pos) = users.iter().position(|u| u.username == username.as_str()) {
        users[pos] = new_user.into_inner();
        HttpResponse::Ok().json("Success")
    } else {
        HttpResponse::NotFound().finish()
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let state = web::Data::new(AppState {
        users: Mutex::new(Vec::new()),
    });

    HttpServer::new(move || {
        App::new()
            .app_data(state.clone())
            .route("/ping", web::get().to(ping))
            .route("/users", web::get().to(get_users))
            .route("/users", web::post().to(post_user))
            .route("/users/{username}", web::get().to(get_user))
            .route("/users/{username}", web::delete().to(delete_user))
            .route("/users/{username}", web::patch().to(patch_user))
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}