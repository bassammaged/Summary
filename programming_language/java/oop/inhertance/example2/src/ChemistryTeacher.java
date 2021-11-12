class ChemistryTeacher extends Teacher{

    public String mainSubject = "Chemistry";

    ChemistryTeacher(){
        this.setCollegeName("Black Toppers");
        System.out.println("Chemistry Teaching.");
        
        System.out.println(getCollegeName());
    }
}