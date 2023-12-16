
import { StyleSheet, Text, Pressable, TouchableOpacity} from "react-native"

const RegisterButton = (props) => {
    const {color, text, press} = props

    return (
        <TouchableOpacity style={styles.pressable} activeOpacity={0.8} onPress={press}>
            <Text style={styles.text}>{text}</Text>
        </TouchableOpacity>
        
    )
}

const styles = StyleSheet.create({
    pressable: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        borderRadius: 10,
        marginTop: 15,
        marginBottom: 10,
        width: 250,
        height: 45,
        backgroundColor: "#015CB8"
    },
    text: {
        color: "white"
    }
})

export default RegisterButton