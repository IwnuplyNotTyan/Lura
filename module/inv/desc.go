package inv

func getEffectDescription(effect string) string {
    switch effect {
    case "heal":
        return "Heals"
    case "damage_boost":
        return "Boosts damage by"
    case "stamina_restore":
        return "Restores stamina by"
    default:
        return effect
    }
}

func getEffectDescriptionUA(effect string) string {
    switch effect {
    case "heal":
        return "Лікує"
    case "damage_boost":
        return "Збільшує пошкодження на"
    case "stamina_restore":
        return "Відновлює витривалість на"
    default:
        return effect
    }
}

func getEffectDescriptionBE(effect string) string {
    switch effect {
    case "heal":
        return "Лякуе"
    case "damage_boost":
        return "Павялічвае пашкоджанні на"
    case "stamina_restore":
        return "Аднаўляе вынослівасць на"
    default:
        return effect
    }
}
