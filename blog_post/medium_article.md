
# A Case for Open Standards

Open standards is a big issue in the software world. I am not an expert on the topic but that doesn’t make sensing benefits of it less apparent for me and I am going to share my recent development experience that is tied to two differing standards to rest my case for genuine open standards.

### Open Standards 101

If no web-browser ever created fit your needs very-well you could be able to create your very own custom web-browser given enough time, tenacity and manpower (or programmer-power to be more precise). That is possible because Web is an open standard. You could find all the latest specifications for HTML, CSS and JS online and implement these specifications into a web browser of your own liking without hassles of patents, fees etc.

As a result of open standard nature of the Web, different web engines like Blink (used by Chrome, Opera and Brave), Gecko (used by Firefox), Safari (used by Apple) compete for superiority. Although these engines could differ in various aspects like performance metrics, degree of conformity to the latest Web specifications, underlying code quality, they all strive for best possible implementation of Web specifications. For a moment, imagine possible implications of Web that was not based on open standards.

One could see similar themes for many other technologies. For example C18 is the latest ISO standard for C programming language and there exists different compilers like GCC, LLVM Clang that support this standard of the language. Many file formats like PNG, CSV and ePub are open standard file formats. You wouldn’t need a propriety software to view these formats. if needed one could write their own code to view files in these formats.

On the flip side one could see closed formats that are controlled by propriety software. Adobe Flash could fit the bill as an exemplar. It wasn’t an open standard technology. I couldn’t remember any serious competitors to Adobe’s Flash Player from pre-HTML5 days of the Web and it should not sound so surprising since Flash wasn’t an open standard format. As a side note Flash’s distancing from open standards that was highlighted at [Steve Jobs 2010 open letter that explains why they won’t support Flash Player on their mobile devices](https://www.theguardian.com/technology/blog/2010/apr/29/steve-jobs-flash-ipad-letter-dead) is noteworthy.

### Problems with DOCX

If you are not a regular Microsoft Office user you could probably relate. DOCX files are problematic. Sometimes Microsoft Office alternatives like LibreOffice, Google Docs couldn’t open them. Other times they could open these documents but couldn’t display them as same as in Microsoft Word. From my own experiences they have a great potential to render one-page original documents in two pages! As a result, after a while I started showing symptoms of DOCX-phobia. I believe many non-Microsoft Office computer users suffer from this format in one way or another. At first you could put the blame on developers of Microsoft Word competitors for aforementioned problems. You could argue that they are incompetent to implement open standard of Office Open XML (OOXML) of DOCX.

If you just scratch the surface you may find out OOXML is another open standard but if you delve into details things become more blurry. I don’t want to go on to technical details of OOXML of DOCX. There is actually [a great article](https://brattahlid.wordpress.com/2012/05/08/is-docx-really-an-open-standard/) that discuss whether DOCX could be considered an open standard. Spoiler alert: According to article we couldn’t say it is a true open standard. It would be more apt to view DOCX format as semi-open or pseudo-open standard. From articles I sifted through about OOXML’s standardization process I sensed serious contention to its standardization process. Some of the supplementary reading material mentioned at the end of this article touches to this topic.

![Demonstration against OOXML in Oslo, Norway in 2008](https://cdn-images-1.medium.com/max/2560/0*eG3pwv50c19VNSoP.jpg)*Demonstration against OOXML in Oslo, Norway in 2008*

Although I didn’t do a through research on causes of DOCX rendering problems on non-Microsoft Office products, I suspect it is highly related to DOCX’s non-conformity to open standards.

### My Latest Encounter with DOCX

In this section I am going to tell about my latest encounter with DOCX documents. I admit next few paragraphs could be a little boring to read but rest of the article is contingent upon some background information.

It is finally time to tell a little bit about myself. I am working as a teaching assistant in a computer engineering department.

Students in our department needs to complete 60 work-days long practical-work at a company as an intern to be conferred with bachelors degree.

This summer because of the Covid-19 pandemic our department followed the suit and offered its students various projects that can be done from their homes instead of on-site work to complete their internships. Demand for project-based internship was really high. We received over 300 project submissions after the deadline.

Over-demand posed some paperwork related problems for internship commission. Our university asked to fill a poorly prepared Word document from each student and submit this form besides their project files and project report. Students were expected to fill student ID, name, begin and end date of their project fields of this form. Committee members need to fill some other fields of this form and print and sign the form afterwards.

Unfortunately, students didn’t conform to a single standard when they filled relevant fields. Some students preferred to capitalize their surnames but not others. There was variability between font family, font size and text color choices to name a few. Added to that the form was a DOCX document🤦 that doesn’t render properly in LibreOffice. To deal with rendering problems I later asked students to send their forms in PDF besides DOCX but that didn’t help much because most of our students continued to send forms only in DOCX format after the announcment.

### Circumventing DOCX

There are some libraries that generate DOCX documents programmatically like [docx.js](https://docx.js.org/#/usage/paragraph) but I didn’t want to spend hours to figure out how to use a new library just to automate form creation task. Besides I have doubts whether I could reproduce a similar form document programmatically with one of these libraries without much strain even if I learned how to work with them.

Content of all the form documents I needed were essentially same but few text fields. I wondered if I could locate these varied parts of the forms by opening DOCX document with a text editor and search for values like student ID and name written to the document to locate fields that vary across different documents. This attempt failed as soon as I realized that text is not stored in an intelligible (clear text) way in DOCX documents.

After failed attempt to locate varied fields in DOCX form document it stroke on me to try to test same strategy after converting DOCX document into open standard odt (Open Document Text) document on LibreOffice. In a nutshell Open Document format is a file that is formed by zip compression of several files. By the way in this respect odt and DOCX formats are very similar. As I tried to read this odt document as plain text without prior decompression, I failed once again. I should note that I could have succeeded in this step if I were to decompress odt file prior to reading. Shortly after my second failure I discovered odt has an easier to work sibling format called fodt that keeps its data in flat XML — good old intelligible XML in uncompressed way. Bingo!

After this great revelation I opened en empty DOCX form in LibreOffice and converted it to fodt format and then did some editing on it to match it’s appearance with original form document in DOCX that is rendered on Microsoft Office as close as I can. Afterwards, I located text fields that needs to be filled by students to create a template document out of fodt file. Finally I wrote a program that would be able to generate form documents automatically from a CSV file where I recorded all necessary information and more for auto-generating forms. Thanks to LibreOffice’s headless capabilities my program was able to create a single PDF document that comprises all student forms that are needed to be printed and then to be signed by myself.

### ODT ➡️ FODT | DOCX ➡️ ???

I think fodt is a great format for making automation of document generation tasks easier. As far as I know there is no equivalent format for DOCX. I believe an equivalent format would expose Microsoft Office’s implementation of OOXML specification and that kind of exposure probably won’t fit into Microsoft’s agenda well unless a paradigm shift happens regarding their OOXML format.

### Embracing DOCX

Just before finalizing this article I found out that I could have located varied text fields in DOCX form with more hardship than performing same task on fodt form documents. I could then have written a form automation program with a more arduous work. This is not all good news though. Since there is no equivalent headless mode in Microsoft Office, my imaginary automation program wouldn’t be able to generate PDFs from DOCx documents. To refrain rendering problems I would have ended up using a friend’s computer to print out many DOCX documents one by one instead of printing all form documents at once with help of a PDF CLI program such as [pdfunite](http://manpages.ubuntu.com/manpages/bionic/man1/pdfunite.1.html).

### Takeaways

If open standard Open Document Format and its XML based intelligible sibling fodt had never existed, I probably wouldn’t able to write a program without much effort that would automate form generation.

In the end I probably have spent more time writing form generation program than doing menial office work like uploading documents to Google Drive to open them in Google Docs (Google Docs renders our student form documents fine unlike LibreOffice), make some styling for each of them like using same font family, capitalizing surnames and then downloading document to hard drive. Repeating this task around 100 times was scary enough for me, so that I ended up spending more time for automating this task.

![One of my favorite programming meme](https://cdn-images-1.medium.com/max/2000/0*hLdQYr8-douBFXvh.png)*One of my favorite programming meme*

![*Perils of automation *from [XKCD: The General Problem](https://xkcd.com/974/)](https://cdn-images-1.medium.com/max/2000/0*C3hVG9hy22hoiuKK.png)**Perils of automation *from [XKCD: The General Problem](https://xkcd.com/974/)*

I think my choice had paid off in few different ways. First of all I had a hands-on lesson about obvious benefits of open standards and shared my experience by writing this article. Along the way I got some insight about DOCX and Open Document formats and learned about existence of headless capabilities of LibreOffice. If I need to deal with a similar menial task in word processing I am now in a stronger position to tackle the hassle!

### Project Repositories

You could check [my form automation project repository](https://github.com/gusanmaz/InternshipFormGenerator) to learn more about project’s inner workings from its documentation or to adapt it to your own needs.

There is one step I skipped in the article about my attempts to tackle this paperwork problem. Before writing an automation program I created a website that would automatically generate a downloadable PDF form for a particular student after that student’s ID information is obtained. This website subjectively created better looking forms with same content, however since I was not able to get an approval to use these forms in lieu of university’s original DOCX forms. Thus this project was another — the very first actually — failed attempt. You could also check [codes for this site from it’s repository](https://github.com/gusanmaz/WebInternshipForms).

### Supplementary Reading

Along the way I sifted through some articles and forum posts regarding open standards, OOXML and Open Document formats. I compiled a list of sources below that was noteworthy to me during my exploration. If you are interested in finer details and more insight these sources might be helpful.

* [What are open standards? | Opensource.com](https://opensource.com/resources/what-are-open-standards)

* [https://stephesblog.blogs.com/papers/stdsprimer.pdf](https://stephesblog.blogs.com/papers/stdsprimer.pdf)

* [Lost in Translation: Interoperability Issues for Open Standards by Rajiv C. Shah, Jay P. Kesan :: SSRN](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=1201708)

* [Standardization of Office Open XML — Wikipedia](https://en.wikipedia.org/wiki/Standardization_of_Office_Open_XML)

* [IBM VP: Office OpenXML a dead end, Microsoft will back ODF | Ars Technica](https://arstechnica.com/information-technology/2008/08/ibm-vp-office-openxml-a-dead-end-microsoft-will-back-odf/)

* [Microsoft’s Great Besmirching | Linux Journal](https://www.linuxjournal.com/content/microsofts-great-besmirching)

* [MS Office 2007 OOXML file format (docx, xslx, pptx, ppsx) (View topic) • Apache OpenOffice Community Forum](https://forum.openoffice.org/en/forum/viewtopic.php?f=74&t=4542)

* [Office 2007 beginning to Bite (View topic) • Apache OpenOffice Community Forum](https://forum.openoffice.org/en/forum/viewtopic.php?f=52&t=258)

* [Advantages of FODT format in LibreOffice | box.matto.nl](https://box.matto.nl/fodtformat.html)
