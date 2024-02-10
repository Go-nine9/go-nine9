import { useEffect,useState } from 'react'
import { useParams } from 'react-router-dom';
import { MapPin,PhoneCall } from 'lucide-react';
import './salon.css';
const Salon = () => {
    const [salon, setSalon] = useState([]);
    let { id } = useParams();
    useEffect(() => {
        document.title = 'Coiffeur: Les meilleurs coiffeurs à Paris 75000'
        fetch(`http://localhost:8097/salons/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
        })
        .then(data => data.json())
        .then(data => {
            setSalon(data)
        })
        
    }
    , [id]);

    return (
        <div className="salonPage">
            <header>
                <div className="info">
                    <h2>{salon.Name}</h2>
                    <p><MapPin />{salon.Address}</p>
                    <p><PhoneCall />{salon.Phone}</p>
                </div>
                <a href="#prestation" className="btn-primary" >Prendre RDV</a>
            </header>
            <div className="parent">
                <div className="div1"></div>
                <div className="div2"></div>
                <div className="div3"></div>
                <div className="div4"></div>
                <div className="div5"></div>
            </div>
            <h2>Réserver en ligne pour un RDV chez {salon.Name}</h2>
            <p className="description">24h/24 - Paiement sur place - Confirmation par mail</p>
            <section className="prestation" id="prestation">
                <h3>Choix de la prestation</h3>
                <div className="flex-container">
                    <div className="prestation-content">
                        <h4>Coupe - Homme</h4>
                        <p className="description">Chaque prestation comprend un diagnostic où nous prenons le temps d'échanger sur vos envies et attentes, suivi d'un shampoing avec le traditionnel massage crânien (soin non inclus), ensuite nous passons à la coupe et procédons pour finir au brushing/coiffage.</p>
                        <div className="service-container">
                        {Array.from({ length: 5 }).map((_, index) => (
                            <div className="service-content" key={index}>
                                <h5>Styliste (Etudiant ou - 20ans) - Shampoing + Coupe personnalisée + Coiffage</h5>
                                <p>40€</p>
                                <button className='btn-primary'>Choisir</button>
                            </div>
                            ))}
                        </div>
                    </div>
                    <div className="salon-info">
                        <div className="rate-container">
                            <h4>Note Global</h4>
                            <div className="rate-content">
                                <div className="left">
                                   <span>4.9</span>
                                </div>
                                <div className="right">
                                    <ul>
                                        <li>Accueil <strong>4.9</strong></li>
                                        <li>Propreté <strong>4.9</strong></li>
                                        <li>Cadre & Ambiance <strong>4.9</strong></li>
                                        <li>Qualité de la prestation <strong>4.9</strong></li>
                                    </ul>
                                    <p>1006 clients ont donné leur avis</p>
                                </div>
                            </div>
                        </div>

                        <h4>Horaires d'ouverture</h4>
                        <div className="horaire-container">

                        </div>
                    </div>
                </div>
            </section>
        </div>
    )
}

export default Salon