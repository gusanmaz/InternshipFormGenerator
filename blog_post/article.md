#A Case For Open Standards

Open standards is a big issue in software world. 
First of all I am not an expert on the topic but that doesn't make sensing
benefits of it less apparent for me and I have a good story to tell from 
my recent development experiences to rest my case.

## Open Standards 101 

If no web-browser ever created fits your needs very-well you could be able
to create your very own custom web-browser given enough time, tenacity and 
manpower (or programmer-power to be more precise). This is possible because
Web is an open standard. You could find all the latest specification for
HTML, CSS and JS online and implement these specifications into a web browser
of you liking without hassles like patents, fees etc. 

As a result of open standard nature of different Web engines 
like Blink(used by Chrome, Opera and Brave), Gecko(used by Firefox), Safari
(used by Apple) compete for superiority. Although these engines could differ
in various performance metrics, degree of conformity to the latest Web specifications,
underlying code quality they all strive to implement open standard of Web 
in a best possible way. For a moment think about possible implications of 
it if Web was not an open standard.

One could see similar theme for many other technologies. For example C18
is the latest ISO standard for C programming language and there exists
different compiler like GCC, LLVM Clang that support this standard of the
language. Many file formats like PNG, CSV and ePub are open standard file
formats. You would't need a propriety software to view these formats and
if needed you could write your own code to view files in these formats.

On the flip side one could see closed formats that are controlled by propriety
software. Flash Player could fit the bill as a good example. It wasn't an open
standard technology. This situation  I don't remember any serious alternatives to Adobe's
Flash Player from pre-HTML5 days of the Web and it should not sound
surprising since Flash was not an open standard technology. 
As a side note Flash's distancing from open standards is highlighted at
[Steve Jobs 2010 open letter that explains why they won't support Flash Player
on their mobile devices](https://www.theguardian.com/technology/blog/2010/apr/29/steve-jobs-flash-ipad-letter-dead).

## Problems with Docx

If you are not a regular Windows and Microsoft Office user you could probably relate.
Docx files are problematic. Sometimes Microsoft Office alternatives like
Libre Office, Google Docs couldn't open them. Other times they could open 
these documents but couldn't display them as same as in Microsoft Word.
They have a great potential to render one-page original documents in two pages.
As a result after a while I started showing symptoms of docx-phobia and believe
many non-Microsoft Office computer users suffer from this format.
At first you could put the blame on developers of other office programs for aforementioned problems.
You could argue that they could be incompetent to implement open standard 
of OpenXML of docx. 

If you scratch the surface you may find out OpenXML is another open standard
but if you delve into details things become more blurry. I don't want to go
on to technical details of OpenXML format of docx. There is actually [a great 
article](https://brattahlid.wordpress.com/2012/05/08/is-docx-really-an-open-standard/) that discuss whether docx could be considered an open standard.
Spoiler alert: According to article we couldn't say it could not be considered as
an open standard. It would be more apt to view docx format as semi-open or
pseudo-open standard. Although I didn't do a through research on docx rendering
problems on non-Microsoft Office products, I highly suspecy it is highly related
to docx's non-conformity to open standards.

## My Latest Encounter with Docx Documents

In this section I am going to tell the story of my latest encounter with docx documents.
Reading this section could be a little boring but understanding rest of the article is contingent upon
background information described in this section.

Let me tell you a bit about myself. I am working as a teaching assistant and
tudents at our department needs to complete 60 work-days long practical-work 
at a company as an intern to be conferred with bachelors degree.

This summer because of the Covid-19 pandemic our department followed 
the suit offered our students projects which can be done from their 
homes instead of on-site work. Demand for project-based internship was really. 
We received over 300 project submissions after the deadline.

Over-demand posed some paperwork related problems for internship commission. 
Our university asked to fill a poorly prepared Word document from each 
student and submit this form besides their project files and project report. 
Students are asked to fill student ID, name, begin date of their project 
and their of their project fields of this form. Committee members 
need to fill some other areas of this form and print and sign the form 
afterwards.

Unfortunately students didn't conform to a single standard when they 
filled this fields. Some students to choose to capitalize their surnames 
but not everyone. There was variability between font family, font size and 
text color choices. Added to that the form is in docx Word document 
:man_facepalming: that doesn't render properly in LibreOffice. To deal with
rendering problems I later asked students to send their forms in PDF besides
docx but that didn't help much because most of our students still continued
to send forms only in docx format.

## Circumventing docx

There are some libraries that generate docx documents like 
[docx.js](https://docx.js.org/#/usage/paragraph) but I didn't want to spend
hours to learn how to use a new library just to automate form creation task. `


Content of all the form documents I needed were actually same except few text fields. I wondered
if I could locate these varied parts of the forms by opening docx document with
a text editor and search values like student ID and name written to document to
locate these fields that vary across different document. This attempt failed
as I understand text is not stored in an intelligible (clear text) way in docx documents. 
`
After failed attempt to locate varied fields in docx form document it stroke on
me to try to do test same thing after converting docx document into open standard
odt document on Libre Office and that was another failure. In a nut shell 
OpenDocument format (odt) is a unintelligible text file that is formed by zip
compression of several files. Short after my second failure I discovered odt has a
easier to work sibling called fodt that keeps its data in good old intelligible 
XML format. Bingo!

After this great revelation I opened en empty docx form in LibreOffice and 
after converting it to fodt format did some editing to the it 
to match it's appearance as similar as possible to the original form document 
that is rendered on Microsoft Office. Afterwards I located text fields that 
needs to be filled by students to create a template document out of fodt file.
Finally I wrote a program that would be able to generate form documents automatically
from a CSV file where I recorded all necessary information and more for 
auto-generating forms. Thanks to Libre Office's headless capabilities my program
was able to create a single PDF document that comprises of all forms that are needed to
be printed to to be signed by myself. 

## Takeaways

If open standard OpenDocument Format and its xml based intelligible sibling fodt never 
existed, I probably wouldn't able to write a program without much effort that would 
automate form generation.
 
In the end I probably have spent more time writing form generation program than doing 
menial office work like uploading document to Google Drive to open them in Google Docs, 
make some styling for each form like using same font family, capitalizing surnames and 
then downloading document to hard drive. Repeating this task around 100 times was scary
enough for me so that I ended up spending more time for automating this task.

I think my choice had paid off in few different ways. First of all I had a hands-on lesson
about obvious benefits of open standards and shared my experience by writing this article. 
Along the way I got some insight about docx and Open Document formats learned about 
existence of headless capabilities of LibreOffice. If I need to deal with a similar menial
task in word processing I am now in a stronger position to tackle the hassle!  

## Project Repositories 

You could check [my form automation project repository](https://github.com/gusanmaz/InternshipFormGenerator) to learn more about
project's inner workings from its documentation or to adapt it to your own
needs. 

![Automation Meme](automation_meme.png)

There is one step I skipped in the article about my attempts to tackle this paperwork
problem. Before writing an automation program I created a website that would automatically generate
a downloadable PDF form for a particular student after that student's ID information is inputted.
This website subjectively created better looking forms with same content, however since I was not able
to get an approval to use these forms in lieu of university's docx forms. Thus this project was another
failed attempt. You could also check [codes for this site from it's repository](https://github.com/gusanmaz/WebInternshipForms).      

    


  

   
   
    