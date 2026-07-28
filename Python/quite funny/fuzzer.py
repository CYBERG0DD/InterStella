import urllib.request
import urllib.error
import time

WORDLIST = [
    "admin", "login", "uploads", "backup", "api", 
    "test", "config", "dev", "images", "CVS"
]

def fuzz_target(base_url):
    print(f"[*] Starting Fuzzer against: {base_url}")
    print("[*] ----------------------------------")
    
    if not base_url.endswith('/'):
        base_url += '/'

    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/113.0.0.0 Safari/537.36'
    }

    for word in WORDLIST:
        target_url = base_url + word
        
        req = urllib.request.Request(target_url, headers=headers)
        
        try:

            response = urllib.request.urlopen(req, timeout=3)
            print(f"[+] FOUND (200 OK): {target_url}")
            
        except urllib.error.HTTPError as e:
            
            if e.code == 403:
                print(f"[!] FORBIDDEN (403): {target_url} - (It exists, but access is denied!)")
            elif e.code == 500:
                print(f"[-] SERVER ERROR (500): {target_url} - (We broke something!)")
                
        except urllib.error.URLError:
            print(f"[-] Network Error: Could not connect to {base_url}")
            break 
            
        time.sleep(0.5)

if __name__ == "__main__":
    target = "http://testphp.vulnweb.com"
    
    fuzz_target(target)