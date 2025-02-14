import React, { useState } from "react";
import axios from "axios";
import "./Game.css";

const Game = () => {
  const [chute, setChute] = useState("");
  const [mensagem, setMensagem] = useState("");
  const [tentativas, setTentativas] = useState([]);
  const [acertou, setAcertou] = useState(false);

  const handleChange = (e) => {
    setChute(e.target.value);
  };

  const enviarChute = async () => {
    if (!chute || isNaN(chute)) {
      setMensagem("⚠️ Digite um número válido!");
      return;
    }

    try {
      const response = await axios.post("http://localhost:8080/chute", {
        numero: parseInt(chute),
      });

      setMensagem(response.data.mensagem);
      setTentativas(response.data.historico);
      setAcertou(response.data.acertou);
      setChute("");
    } catch (error) {
      setMensagem("❌ Erro ao enviar o chute. Tente novamente.");
    }
  };

  return (
    <div className="game-container">
      <div className="game-screen">
        <h1>🎯 Jogo da Adivinhação</h1>
        <p>Digite um número entre 0 e 100:</p>
        <div className="input-container">
          <input
            type="number"
            value={chute}
            onChange={handleChange}
            disabled={acertou}
          />
          <button onClick={enviarChute} disabled={acertou}>
            Enviar Chute
          </button>
        </div>
        <p className="mensagem">{mensagem}</p>

        {tentativas.length > 0 && (
          <div className="tentativas">
            <h3>Tentativas:</h3>
            <p>{tentativas.join(", ")}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default Game;
