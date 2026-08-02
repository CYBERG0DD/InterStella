import socket
import sys

def grab_banner(ip, port):
    print(f"[*] Attempting to connect to {ip} on port {port}...")
    
    try:
        
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        
        s.settimeout(3)
        s.connect((ip, port))
        
        banner = s.recv(1024)
        
        print(f"[+] SUCCESS! Port {port} is OPEN.")
        print(f"[+] Banner received: {banner.decode().strip()}")
        
    except socket.timeout:
        print(f"[-] Connection timed out. The port might be filtered by a firewall.")
    except ConnectionRefusedError:
        print(f"[-] Connection refused. The port is definitely closed.")
    except Exception as e:
        print(f"[-] An error occurred: {e}")
    finally:
        s.close()

if __name__ == "__main__":
    target_ip = "127.0.0.1" 
    target_port = 22 
    
    grab_banner(target_ip, target_port)