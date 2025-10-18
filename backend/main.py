from fastapi import FastAPI
import json
from scrapers.remotar import buscar_remotar

app = FastAPI(title="JobFinder API")

@app.get("/vagas")
def listar_vagas(termo: str = "help desk"):
    vagas = buscar_remotar(termo)
    data = {"termo": termo, "total": len(vagas), "vagas": vagas}
    with open("data/vagas.json", "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    return data
