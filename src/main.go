package main

// std
import(
	"fmt"
)

// internal
import(
	"learn-design-patterns-GoF/patterns/creational/abstractfactory"
)


func main() {
	// ################################
	//  Design Pattern Abstract Factory => more information in https://refactoring.guru/design-patterns/abstract-factory
	fmt.Println("\n\nDesign Pattern Abstract Factory")
	
	cat := abstractfactory.NewCat("Stewie", "Cat")
	abstractFactory1 := abstractfactory.NewAbstactFactory(cat)
	fmt.Println("\n" + abstractFactory1.GetAnimalInfo())
  
	// ################################
  
  
	// // ################################
	// // Design Pattern Builder => more information in https://refactoring.guru/design-patterns/builder
	// echo "\n\nDesign Pattern Builder\n"
  
	// let setup = newSetup()
	// let admin = newAdmin[Setup](setup)
	// var car: Car = admin.build()
	// echo car.getCarInfo()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Factory Method => more information in https://refactoring.guru/design-patterns/factory-method
	// echo "\n\nDesign Pattern Factory Method\n"
  
	// let creator = newCreator()
	// var company: Company1 = create[Company1](creator)
	// echo company.getCompanyInfo()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Prototype => more information in https://refactoring.guru/design-patterns/prototype
	// echo "\n\nDesign Pattern Prototype\n"
  
	// var zaraClothing = newZara("Zara", "Jeans", 19.99, "Mafia Guy")
	// echo zaraClothing.getClothingInfo()
  
	// var zaraClothing2 = zaraClothing.copy()
	// echo zaraClothing2.getClothingInfo()
  
	// zaraClothing2.setModel("T-Shirt")
	// zaraClothing2.setPrice(15.99)
  
	// echo zaraClothing.getClothingInfo()
	// echo zaraClothing2.getClothingInfo()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Singleton => more information in https://refactoring.guru/design-patterns/singleton
	// echo "\n\nDesign Pattern Singleton\n"
  
	// getInstance().addCar("VW Golf")
	// getInstance().addCar("VW Passat")
	// getInstance().addCar("Ford Fiesta")
	// echo getInstance().getCars()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Adapter => more information in https://refactoring.guru/design-patterns/adapter
	// echo "\n\nDesign Pattern Adapter\n"
  
	// echo "Object Adapter\n"
  
	// let government = newGovernment()
	// var governmentImpl = newGovernmentAdapterImpl(government)
	// var broker = newBroker[GovernmentAdapterImpl](governmentImpl)
  
	// echo broker.sendGains(parseJson("""{"taxPayerId":1234567890, "taxPayerGains":10000}"""))
  
  
	// echo "\nClass Adapter\n"
  
	// let classAdapter = newClassAdapter()
	// var broker2 = newBroker[ClassAdapter](classAdapter)
  
	// echo broker2.sendGains(parseJson("""{"taxPayerId":1234562490, "taxPayerGains":156000}"""))
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Bridge => more information in https://refactoring.guru/design-patterns/bridge
	// echo "\n\nDesign Pattern Bridge\n"
  
	// var tv = newTv()
	// var radio = newRadio()
  
	// var remoteRadio = newRemote[Radio](radio)
	// echo "Radio"
	// echo radio.getDeviceInfo()
	
	// discard remoteRadio.volumeUp()
	// discard remoteRadio.tooglePower()
  
	// echo radio.getDeviceInfo()
  
	// var remoteTv = newRemote[Tv](tv)
	// echo "\nTV"
	// echo tv.getDeviceInfo()
  
	// discard remoteTv.volumeDown()
	// discard remoteTv.channelUp()
  
	// echo tv.getDeviceInfo()
  
  
	// var advancedRemoteTv = newAdvancedRemote[Tv](tv)
  
	// echo "\nTV with Advanced Remote"
	// echo tv.getDeviceInfo()
  
	// discard advancedRemoteTv.mute()
  
	// echo tv.getDeviceInfo()
	
	// // ################################
  
  
	// // ################################
	// // Design Pattern Composite => more information in https://refactoring.guru/design-patterns/composite
	// echo "\n\nDesign Pattern Composite\n"
  
	// var circle = newCircle(0, 0, PI)
	// var dot = newDot(0, 0)
	
	// var coumpoundGraphic = newCoumpoundGraphic()
	// coumpoundGraphic.add(circle)
	// coumpoundGraphic.add(dot)
  
	// discard coumpoundGraphic.draw()
	// discard coumpoundGraphic.move(1,1)
	// discard coumpoundGraphic.draw()
  
	// // ################################
}