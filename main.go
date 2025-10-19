package main

import "fmt"

//Création des structures de contact
type contact struct {
	ID    int
	name  string
	email string
	phone string
}
type contactmanager struct {
	contacts []contact
	nextID   int
}

//méthode pour créer un contact et le mettre dans la liste
func (cm *contactmanager) addcontact(name, email, phone string) {

	//nouveau contact avec id
	X := contact{
		ID:    cm.nextID,
		name:  name,
		email: email,
		phone: phone,
	}

	cm.contacts = append(cm.contacts, X)
	cm.nextID++
	fmt.Printf("contact au nom de %s ajouté à la base!\n", X.name)
	fmt.Printf("identifiant attribué: %d\n", X.ID)
}
func (cm *contactmanager) showcontact() {
	fmt.Println("---voici la liste des contacts très cher!!---")

	for _, X := range cm.contacts {
		fmt.Printf("ID: %d|| nom: %s|| email: %s|| téléphone: %s\n", X.ID, X.name, X.email, X.phone)
	}

}

//recherche de contact existant dans la liste (recherche par id)
func (cm *contactmanager) getcontact(id int) (contact, error) {
	fmt.Printf("---votre recherche pour l'identifiant %d---\n", id)

	for _, X := range cm.contacts {
		if X.ID == id {

			return X, nil
		}
	}
	return contact{}, fmt.Errorf("aucun contact trouvé avec l'id: %d", id)
}

//recherche un ou plusieurs contacts existants dans la liste (recherche par nom)
func (cm *contactmanager) Searchbyname(query string) []contact {
	fmt.Printf("---votre recherche pour le nom %s---\n", query)
	var resultat []contact

	for _, X := range cm.contacts {
		if X.name == query {
			resultat = append(resultat, X)

		}

	}
	if len(resultat) == 0 {
		fmt.Printf("aucun resultat pour le nom: %s\n", query)
	} else {
		fmt.Printf("---voici tous les contacts qui contiennent le nom: %s---\n", query)
	}
	return resultat

}

//Recherche un contact par email
func (cm *contactmanager) Searchbyemail(email string) (contact, error) {

	for _, X := range cm.contacts {
		if X.email == email {
			return X, nil
		}
	}
	return contact{}, fmt.Errorf("Aucun contact trouvé avec l'Email: %s", email)
}

//affichage de la liste complète de contact

func (cm *contactmanager) listcontact() []contact {
	return cm.contacts
}

//Supprimer avec confirmation un contact s'il existe avec Soit le nom, ou l'ID
func (cm *contactmanager) deletecontact() error {
	var choix int
	fmt.Println("voulez-vous supprimer le contact par:")
	fmt.Println("1. ID")
	fmt.Println("2. Nom")
	fmt.Print("votre choix: ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		var id int
		fmt.Print("Entrez l'ID à supprimer: ")
		fmt.Scan(&id)

		for i, X := range cm.contacts {
			if X.ID == id {
				var confirmation string
				fmt.Printf("voulez-vous vraiment supprimer %s (ID: %d) ? (o/n):", X.name, X.ID)
				fmt.Scan(&confirmation)

				if confirmation == "o" || confirmation == "O" {
					nom := X.name
					cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
					fmt.Printf("l'id: %d appartenant à %s est supprimé de la liste de contact\n", id, nom)
					return nil

				} else {
					fmt.Println("suppression annulée")
					return nil
				}

			}
		}
		return fmt.Errorf("Aucun contact trouvé avec l'ID: %d", id)
	case 2:
		var nom string
		fmt.Print("Entrez le nom à supprimer: ")
		fmt.Scan(&nom)

		for i, X := range cm.contacts {
			if X.name == nom {
				var confirmation string
				fmt.Printf("voulez-vous vraiment supprimer %s (ID: %d) ? (o/n):", X.name, X.ID)
				fmt.Scan(&confirmation)

				if confirmation == "o" || confirmation == "O" {
					nom := X.name
					id := X.ID
					cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
					fmt.Printf("l'id: %d appartenant à %s est supprimé de la liste de contact\n", id, nom)
					return nil

				} else {
					fmt.Println("suppression annulée")
					return nil
				}

			}
		}
		return fmt.Errorf("aucun contact trouvé avec le nom %s", nom)

	}

	return fmt.Errorf("choix invalide")
}

func main() {
	cm := &contactmanager{}
	cm.nextID = 1

	for {
		fmt.Println("\n=== Gestionniaire de Contacts ===")
		fmt.Println("1.ajouter un contact")
		fmt.Println("2.lister un contact")
		fmt.Println("3.rechercher un contact")
		fmt.Println("4.recherche de doublons de contacts")
		fmt.Println("5.supprimer un contact")
		fmt.Println("6.Quitter")

		var choix int
		fmt.Print("veuillez entrer une commande:")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var name, email, phone string
			fmt.Print("nom:")
			fmt.Scan(&name)
			fmt.Print("email:")
			fmt.Scan(&email)
			fmt.Print("téléphone:")
			fmt.Scan(&phone)
			cm.addcontact(name, email, phone)

		case 2:
			cm.showcontact()

		case 3:
			var choice int
			fmt.Printf("Voulez-vous rechercher par:\n")
			fmt.Print("1.Email:\n")
			fmt.Print("2.Id:\n")
			fmt.Println("faites votre choix: ")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				var email string
				fmt.Printf("Entrez l'Email du contact à rechercher:")
				fmt.Scan(&email)
				c, err := cm.Searchbyemail(email)
				if err != nil {
					fmt.Printf("Aucun Contact trouvé avec l'Email: %s", email)

				} else {
					fmt.Printf("----Votre recherche pour l'email: %s----\n", email)
					fmt.Printf("ID %d || nom: %s || email: %s || téléphone: %s\n", c.ID, c.name, c.email, c.phone)
				}
			case 2:
				var id int
				fmt.Printf("entrez l'id du contact à rechercher:")
				fmt.Scan(&id)
				c, err := cm.getcontact(id)
				if err != nil {
					fmt.Println("contact non trouvé!!!")
				} else {
					fmt.Printf("contact trouvé: ID %d || nom: %s || email: %s || téléphone: %s\n", c.ID, c.name, c.email, c.phone)

				}
			default:
				fmt.Printf("Choix invalide!\n")
			}
		case 4:
			var query string
			fmt.Println("entrez le nom du contact(s) à rechercher")
			fmt.Scan(&query)
			resultat := cm.Searchbyname(query)
			if len(resultat) == 0 {
				fmt.Printf("Aucun contact(s) trouvé au nom de %s\n", query)

			} else {
				for _, X := range resultat {
					fmt.Printf("resultat trouvé au nom de %s\n", query)
					fmt.Printf("ID: %d || Nom: %s || email: %s || tel: %s\n", X.ID, X.name, X.email, X.phone)

				}
			}
		case 5:
			err := cm.deletecontact()
			if err != nil {
				fmt.Println(err)
			}
		case 6:
			fmt.Println("Merci d'avoir recommandé notre gestionnaire")
			fmt.Println("Gardez la pêche")
			return
		default:
			fmt.Println("choix invalide, essayez encore.")

		}

	}

}
