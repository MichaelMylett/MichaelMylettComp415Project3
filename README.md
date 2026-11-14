# MichaelMylettComp415Project3

For this app, it is important that you have the window in fullscreen to properly see all of the information. It is also important to mention that the program uses an API key saved under public variable APIKey in a file which will *NOT* be attached in the GitHub, for security purposes. For the program to work, you must define APIKey in one of the files provided, or add a new file in the same package containing the APIKey variable and definition.

This project uses the Fyne library to create a GUI for a user to find information on countries through. The user is given a dropdown of several regions and several empty boxes on the right side of the screen (some with placeholder text). These boxes will later be filled with the official name, common name, official languages, national flag, landlocked status, population, currencies, and region of a country of choice. I felt that letting the user select by region would be simpler to understand (as the regions are based on the continents) and simpler for me to code.

When the user selects one of the regions from the dropdown (this excludes Antarctica, as there are of course no countries there), they will be given a list of up to 100 countries in that region (which I believe to be more than enough to contain every country in any of the regions). The user can then click on a country, and all the previously mentioned boxes on the righthand side will be filled with the correct information/imagery. If the user clicks on any other country, the information will all update to that of the new country selected. This works still if switching regions between each selection of a country.

The only question anybody asked me about my program was how it was getting the data. I explained the idea of a RestAPI as essentially being an online database setup with features to allow me to interact with it via an API key, and how the data was being pulled from the internet rather than it being stored somewhere in my computer. Other than that, I did not get any real questions as to how it works.

I was able to complete every part of the assignment.

Most of the structs for pulling from RestCountries were made by AI, and for several other parts of the project I used the RestDemo from class as a template of sorts. I did need to look up many commands from the Fyne library to get the project to work (especially for the flag image), but other than the structs I did not copy much from AI.
