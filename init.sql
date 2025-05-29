-- Create products table
CREATE TABLE IF NOT EXISTS produits (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    prix DECIMAL(10,2) NOT NULL,
    categorie VARCHAR(50) NOT NULL,
    quantite INTEGER NOT NULL DEFAULT 0
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    caisse VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('VENTE', 'RETOUR')),
    produit_ids VARCHAR(100) NOT NULL,
    montant DECIMAL(10,2) NOT NULL,
    deja_retourne BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create caisse table
CREATE TABLE IF NOT EXISTS caisses (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    occupe BOOLEAN NOT NULL DEFAULT FALSE
);

-- Insert caisses
INSERT INTO caisses (nom, occupe) VALUES
('Caisse 1', FALSE),
('Caisse 2', FALSE),
('Caisse 3', FALSE);

-- Insert sample products
INSERT INTO produits (nom, prix, categorie, quantite) VALUES

-- Boissons
('Eau minérale 1.5L', 1.20, 'Boissons', 50),
('Coca-Cola 2L', 2.50, 'Boissons', 30),
('Jus d''orange 1L', 2.80, 'Boissons', 25),
('Café en grains 250g', 4.50, 'Boissons', 20),
('Thé vert 100 sachets', 3.90, 'Boissons', 15),

-- Produits laitiers
('Lait entier 1L', 1.30, 'Produits laitiers', 40),
('Yaourt nature 4x125g', 2.20, 'Produits laitiers', 35),
('Fromage râpé 200g', 2.80, 'Produits laitiers', 25),
('Beurre doux 250g', 2.50, 'Produits laitiers', 30),
('Crème fraîche 20cl', 1.80, 'Produits laitiers', 20),

-- Épicerie
('Pâtes spaghetti 500g', 1.20, 'Épicerie', 45),
('Riz basmati 1kg', 2.30, 'Épicerie', 30),
('Huile d''olive 1L', 5.90, 'Épicerie', 20),
('Sucre en poudre 1kg', 1.50, 'Épicerie', 40),
('Sel fin 1kg', 0.90, 'Épicerie', 50),

-- Snacks
('Chips nature 150g', 1.80, 'Snacks', 60),
('Biscuits chocolat 300g', 2.40, 'Snacks', 40),
('Barres chocolatées x5', 3.20, 'Snacks', 35),
('Pop-corn 200g', 2.10, 'Snacks', 25),
('Crackers 200g', 1.90, 'Snacks', 30),

-- Hygiène
('Dentifrice 100ml', 2.30, 'Hygiène', 25),
('Savon liquide 300ml', 3.50, 'Hygiène', 20),
('Shampoing 400ml', 4.20, 'Hygiène', 15),
('Déodorant 150ml', 3.80, 'Hygiène', 20),
('Papier toilette 8 rouleaux', 2.90, 'Hygiène', 30); 