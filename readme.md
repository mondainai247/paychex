#Paychex Ignite CLI Tutorial

Hello Everyone, let me show you a simple tutorial that shows the simplicity and power of using Ignite CLI to create powerful Application specific blockchains.
For this demonstration, we will create an application for doing a company's payroll. I assume you already have ingite installed, if you don't please click here for my other tutorial on installing ignite CLI. 
[Install Ignite CLI Tutorial](https://medium.com/@himitsu1007/installing-ignite-cli-on-linux-cdf77a5436f3)


Requirements:
1) An employee roster 
2) A way to classify employees by their Role in the company
3) A way to ensure the employees can be paid only once per month.

Okay, let's get started, let's scaffold a chain called paychex, and let's change the address prefix to paychex.

```
#open your command prompt. 
cd #into your home directory
ignite scaffold chain paychex --address-prefix paychex
```
If done correctly, you should see something like this:
![scaffold paychex](/paychex/docs/images/scaffold paychex.png)

If you get an error message, it is most likely that you do not have go enabled in your operating environment, you can fix that with this command.
```
export GO111MODULE = on
```
Let's change our current working directory to our newly scaffolded project directory.
```
cd paychex
```
Now in order for us create our application, we need to create some custom types. Think of these as custom variables that are effectively (multi)variable objects/structs. We can think of these like a new type of data structure, similar to how stings, bools are unique.

We can create these with the ignite scaffold type command(Or list, Map commands if we want CRUD functionality). This will require us to name the custom type in this case it will be "Role" and we will want to have at least one parameter to add, in this came "name", so in order to scaffold that we would use the following command:
```
ignite scaffold type Role name
```
![scaffold role](paychex/docs/images/scaffold role.png)

This created a new proto file in the proto/paychex/paychex directory. 
Now lets call that Type when scaffolding a List. Now let's create an employee type Employee. 
```
ignite scaffold list Employee name role:Role 
```
![scaffold Employee](paychex/docs/images/scaffold Employee List.png)
Notice that we pass the Role to the "role" parameter. 
Be sure to click "y" for yes. 
![Employee fiels](/paychex/docs/images/scafollded Employee files.png)
Wow! Lots of files have been scaffolded by ignite CLI. Sure glad we didnt have to type those all by ourselves!
Let's check the proto/paychex/paychex/employee.proto file to see its format.```
nano proto/paychex/paychex/employee.proto
```
![Employee Proto](paychex/docs/images/Employee Proto.png)
See how the role parameter as automatically taken our custom "Role" Type? By scaffolding Employee as a list, it will add two more parameters, id & creator. 
Let's create a way to Pay the Employees. 
For that we will need a type called Payroll and a custom message called send-payroll. 
```
ignite scaffold type Payroll year:int month:int employees:uint
```
![Payroll Proto](paychex/docs/images/Payroll proto.png)

Open the payroll proto file and add the word repeated to the proto file. 
We are also going to have some custom queries. For that we need to scaffold some types. 
```
ignite scaffold type Staff roleName employeeCount:uint
```
Next we will scaffold a List type, QueryResponse:

```
ignite scaffold list QueryResponse staff:Staff
```
![Query Response Files](paychex/docs/images/Query Response Files.png)
We can pass it the staff. Note: it would have been my preference to scaffold this as a Type and not list, however the blockchain will not launch as their will be no staff.go file created. (when I figure out a soltion, I will update this accordingly). 
![Query Response Proto](paychex/docs/images/query response proto.png)

We need to open the proto/paychex/paychex/query_response.proto file and add the word "repeated" before the Type staff. 
Before we launch our chain. Let's add two more custom message types. 
```
ignite scaffold message send-payroll newpayroll:Payroll  
```
![scaffold payroll](paychex/docs/images/scaffold send payroll.png)
```
ignite scaffold query list-staff - response QueryResponse
```
We are ready to launch our chain. 
```
cd 
cd paychex
ignite chain serve
```
If everything was done correctly, a Blockchain will launch. Please click allow, otherwise your chain will not accept messages and or queries. It will be only open to your local host(not the outside word).
![ignite chain serve](paychex/docs/images/ignite chain serve.png)
We can see an App binary was created in our ~/go/bin folder
We can see our functionality for the binary with this command. 
```
paychexd --help
```
![paychex help](paychex/docs/images/paychexd help.png)
If you don't see this, then you need to make sure your go/bin is on your path. 
Let's check and see if we have our create-employee tx funtion. This was automatically generated because we made "Employee" a List.
```
paychexd tx paychex --help
```
![paychex tx help](paychex/docs/images/paychexd help.png)
There it is at the top! 
If we input the tx and add the - help flag we can see the parameters. 
![paychex send-payroll help](paychex/docs/images/paychexd tx paychex send-payroll.png)
We can create a new employee with the command below. (Note: the syntax needs to be correct or the transaction will fail). 
```
paychexd tx paychex create-employee "Jose Gonzales" '{"name": "Engineer"}' --from alice
```
![tx create omployee](paychex/docs/images/tx creat-employee.png)

Let's go ahead and add some more staff to our company's payroll system! 
```
paychexd tx paychex create-employee "John Smith" '{"name": "Sales"}' --from alice
paychexd tx paychex create-employee "Tomy Wong" '{"name": "Engineer"}' --from alice
paychexd tx paychex create-employee "Tomomi Suzuki" '{"name": "Engineer"}' --from alice
paychexd tx paychex create-employee "Sally Yu" '{"name": "Engineer"}' --from alice
paychexd tx paychex create-employee "Bruce Carter" '{"name": "Marketing"}' --from alice
paychexd tx paychex create-employee "Phil Brewster" '{"name": "Marketing"}' --from alice
```
Now let's track our payroll. We can use the send-payroll message type we scaffolded. 
```
paychexd tx paychex send-payroll '{"year":2023,"month":8,"employees": 0}' --from alice
```
![tx send payrol](paychex/docs/images/tx send-payroll.png>)
Unfortunately, we do not know how to prevent the same payroll from being sent twice! (Working on that). 
We can also list entire staff with this command. 
```
paychexd query paychex list-employee
```
![list-employee](paychex/docs/images/list-employee.png)
We created a check-staff query as well using our custom query types. Let's try that:
```
paychexd query paychex check-staff
```
![check-staff](paychex/docs/images/paychexd query check-staff.png>)
Oh no! The response was empty. We are going to get fired! Why isn't it working? Let's check the query file in the keeper folder. 
```
nano x/pachex/keeper query_check_staff.go
```
![check staff proto](<paychex/docs/images/query check staff proto.png>)
There is our problem, none of the logic has been implemented. Let's see what it is supposed to look like (or something similar). 

As we can see, we need to implement some message handling logic, if we attempt this without knowing what we are doing we will break the chain, so I will update this tutorial after consulting with other professionals. 

In the meantime, we have a way of tracking employees, tracking if payroll was sent by using a blockchain scaffolded with Ignite CLI. Impressive! 
References:
[Ignite CLI Documentaion](https://docs.ignite.com/)
