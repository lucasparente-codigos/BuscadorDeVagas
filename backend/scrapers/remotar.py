import requests
from bs4 import BeautifulSoup

def buscar_remotar(termo: str):
    url = f"https://remotar.com.br/vagas?search={termo.replace(' ', '+')}"
    resp = requests.get(url)
    soup = BeautifulSoup(resp.text, "html.parser")
    vagas = []
    for card in soup.select("a[href*='/job/']"):
        titulo = card.get_text(strip=True)
        link = card["href"]
        vagas.append({"titulo": titulo, "link": link})
    return vagas
