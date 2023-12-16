import { StyleSheet, Text, Pressable} from "react-native"

const LandingButton = ({backgroundColor, text, navigate}) => {
    return (
        <Pressable onPress={navigate} style={[styles.pressable, { backgroundColor }]}>
            <Text style={styles.text}>{text}</Text>
        </Pressable>
        
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
        height: 45
    },
    text: {
        color: "white"
    }
})

export default LandingButton