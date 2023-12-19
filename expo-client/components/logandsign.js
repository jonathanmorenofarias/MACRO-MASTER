import Header from "./Header"
import Arrow from "./NavigateArrow"


const getInitial = (name) => {
    return(
        {
            headerTitle: () => <Header name={name}/>,
            headerLeft: () => <Arrow/>,
            headerLeftContainerStyle: { paddingLeft: 10 },
            headerStyle: {
                height: 100, 
                backgroundColor: "#424242", 
                borderBottomWidth: .25, 
                borderBottomColor: "grey"
              }
          }
    )
}

export { getInitial }