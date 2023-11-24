use std::fs::File;
use std::io::Read;
use warp::Filter;
use serde::{Deserialize, Serialize};

#[derive(Deserialize,Serialize,Debug)]
struct User {
    name: String,
    key: String,
}
fn read_db() -> Result<Vec<User>, std::io::Error> {
    let mut db = File::open("user.db")?;
    let mut users: Vec<User> = Vec::new();
    //db.read_to_end(&mut users)?;
    Ok(users)
}
#[tokio::main]
async fn main() {
    let user_data = match read_db() {
        Err(e) => {
            let _ = File::create("user.db");
        },
        Ok(res) => res, 
    };
    println!("{:?}", Ok(user_data));


    /*
    // GET /
    let home = warp::path::end()
        .map(|| format!("home endpoint"));

    // POST /register
    let register = warp::post()
        .and(warp::path("register"))
        .and(warp::path::param::<u32>())
        .and(warp::body::json())
        .map(|user|, mut user: User| {
            user.name = name
        }
    warp::serve(home)
        .run(([127,0,0,1],6000))
        .await;
    */
}


