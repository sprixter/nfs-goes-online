use std::net::TcpListener;
use std::net::TcpStream;
use std::fs;
use std::io::prelude::*;

fn main() {
    // create 
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();
    for stream in listener.incoming() {
        let stream = stream.unwrap();
        handle_connection(stream);
    }
}

fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer).unwrap();

    let get: &[u8; 16] = b"GET / HTTP/1.1\r\n";

    let (status_line, filename) = if get[11] == 49 && get[13] == 49 {
        if buffer.starts_with(get) {
            ("HTTP/1.1 200 OK", "index.html")
        } else {
            ("HTTP/1.1 404 NOT FOUND", "404.html")
        }
    } else {
        ("HTTP/1.1 400 BAD REQUEST", "400.html")
    };

    let response = match fs::read_to_string(filename) {
        Ok(contents) => format!(
            "{}\r\nContent-Length: {}\r\n\r\n{}",
            status_line,
            contents.len(),
            contents
        ),
        Err(_) => {
            let contents = fs::read_to_string("500.html").unwrap();
            format!(
                "HTTP/1.1 500 INTERNAL SERVER ERROR\r\nContent-Length: {}\r\n\r\n{}",
                contents.len(),
                contents
            )
        }
    };

    stream.write_all(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}