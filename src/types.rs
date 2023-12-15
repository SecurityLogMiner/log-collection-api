use std::fs::File;
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

#[derive(Deserialize)]
pub struct Config {
    #[serde(default = "default_host")]
    pub host: String,
    pub port: u16,
    pub client_origin_url: String,
}

fn default_host() -> String {
    "127.0.0.1".to_string()
}

impl Default for Config {
    fn default() -> Self {
        envy::from_env::<Config>().expect("Provide missing environment variables for Config")
    }
}

#[derive(Serialize)]
pub struct ErrorMessage {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub error_description: Option<String>,
    pub message: String,
}

