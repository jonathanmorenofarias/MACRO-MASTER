import { StyleSheet, Text, View } from "react-native"


const Header = ({name}) => {
    
    return(
        <View style={styles.view}>
            <Text onPress={() =>  navigation.navigate('Landing')} style={styles.text}>{name}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    view: {
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
        height: "100%"
    },
    text: {
        fontWeight: "bold",
        fontSize: 18,
        color: "white"
    }
})

export default Header