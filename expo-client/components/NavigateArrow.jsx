import { useNavigation } from "@react-navigation/native"
import Ionicons from "@expo/vector-icons/Ionicons"

const NavigateArrow = () => {
    const navigation = useNavigation()
    return(
        <Ionicons name="arrow-back-outline" size={ 25 } color="white" onPress={() => navigation.navigate('Landing')} />
    )
}

export default NavigateArrow