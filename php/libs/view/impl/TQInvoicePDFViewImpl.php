<?php

require_once (dirname(__FILE__). '/../PDFView.php');

define('FPDF_FONTPATH', dirname(__FILE__) . '/../../fpdf/font/');
require_once (dirname(__FILE__).  '/../../fpdf/mbfpdf.php');

function s($str) {
    return mb_convert_encoding($str, "SJIS", "UTF-8");
}

class TQInvoicePDFViewImpl {
    private $pdf;
    
    private $startX = 20;
    private $accColorR = 195;
    private $accColorG = 13;
    private $accColorB = 35;

    public function __construct() {
        $this->pdf = $pdf = new MBFPDF('P', 'mm', 'A4');
        $pdf->AddMBFont(GOTHIC ,'SJIS');
        $pdf->AddMBFont(MINCHO ,'SJIS');
        $pdf->AddMBFont(HIRAKAKU_W3 ,'SJIS');        
        $pdf->AddMBFont(HIRAKAKU_W6 ,'SJIS');
        $pdf->SetAutoPageBreak(false);
        $pdf->SetMargins(0, 0, 0, 0);
        $pdf->AddPage('P');
    }

    public function writeToAddr($addr, $name) {
        $pdf = $this->pdf;
        $pdf->SetFont(HIRAKAKU_W3,'', 10);
        $pdf->SetTextColor(0, 0, 0);
        $pdf->SetXY(25.5, 15.5);
        $pdf->MultiCell(65, 5.3, s($addr), 0, 1, 'L', 0);
        
        $pdf->Cell(25.5, 5.3, s(''), 0, 0, 'C', 0);
        $pdf->SetFont(HIRAKAKU_W6,'', 10);
        $pdf->MultiCell(65, 5.3, s($name), 0, 1, 'L', 1); // next line
    }

    public function writeItemTitle($title) {
        $pdf = $this->pdf;
        $pdf->SetDrawColor($this->accColorR, $this->accColorG, $this->accColorB);
        $pdf->SetLineWidth(0.2);
        $pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $pdf->SetX(25.5);
        $pdf->Cell(65, 6, s($title), 1, 2, 'C', 0);
    }

    public function writeDate($date) {
        $pdf = $this->pdf;
        $pdf->SetFont(HIRAKAKU_W3,'', 6.5);
        $pdf->SetTextColor(77, 77, 77);
        $pdf->SetXY(160, 15.5);
        $pdf->Cell(28.5, 5.8, s(date('Y年n月j日', $date)), 0, 2, 'R', 0);
    }

    public function writeMyCompany($name, $addr) {
        $pdf = $this->pdf;
        $pdf->SetFont(HIRAKAKU_W6, '', 8);
        $pdf->SetXY(140.5, 26);
        $pdf->write(6, s($name));
        
        $pdf->SetFont(HIRAKAKU_W3,'', 7);
        $pdf->SetXY(140.5, 32.5);
        $pdf->MultiCell(45, 4.3, s($addr), 0, 1, 'L', 0);
    }

    public function writeTitleMessage() {
        $pdf = $this->pdf;
        $pdf->SetFont(HIRAKAKU_W6,'', 14);
        $pdf->SetXY(25.5, 65);
        $pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $pdf->Cell(159, 10, s('書類送付状'), 'B', 1, 'C', 0);

        $pdf->SetFont(HIRAKAKU_W3,'', 10);
        $pdf->SetXY(25.5, 80);
        $pdf->SetTextColor(0, 0, 0);
        $pdf->MultiCell(159, 6, s("拝啓\n貴社におかれましては、ますますご清栄のことと心よりお慶び申し上げます。\n平素は格別のご高配を賜り、厚く御礼申し上げます。\n早速ではございますが下記の書類をお送りします。ご査収の上よろしくご手配を賜りますようお願い申し上げます。"), 0, 1, 'L', 0);
        $pdf->SetX(178);
        $pdf->MultiCell(159, 6, s('敬具'), 0, 1, 'R'); // next line

        $pdf->SetX(40);
        $pdf->SetDrawColor(0, 0, 0);
        $pdf->SetFont(HIRAKAKU_W6,'', 10);
        $pdf->Cell(130, 8, s('記'), 'B', 1, 'C'); // next line
    }

    public function writeItems($list) {
        $pdf = $this->pdf;
        $pdf->SetX(60);
        $pdf->SetTextColor(0, 0, 0);
        $pdf->SetFont(HIRAKAKU_W3,'', 10);
        foreach($list as $item) {
            $pdf->Cell(60, 8, s($item['name']), 0, 0, 'L');
            $pdf->Cell(30, 8, s($item['num']), 0, 1, 'R'); // next line
        }
    }

    public function output($name) {
        $pdf = $this->pdf;
        $pdf->SetX(178);
        $pdf->MultiCell(159, 6, s('以上'), 0, 1, 'R'); // next line
        
        $pdf->Output($name . '.pdf', 'D');
    }
}
?>