<?php

require_once (dirname(__FILE__). '/../PDFView.php');

define('FPDF_FONTPATH', dirname(__FILE__) . '/../../fpdf/font/');
require_once (dirname(__FILE__).  '/../../fpdf/mbfpdf.php');

function s($str) {
    return mb_convert_encoding($str, "SJIS", "UTF-8");
}

function getTax($v, $taxRate) {
    return ceil($v * $taxRate / 100);
}

function removeTax($v, $taxRate) {
    return floor($v * 100 / (100 + $taxRate));
}

class PDFViewImpl implements PDFView {
    private $pdf;

    public function __construct() {
        $this->pdf = new MBFPDF('P', 'mm', 'A4');
        $this->pdf->AddMBFont(GOTHIC ,'SJIS');
        $this->pdf->AddMBFont(MINCHO ,'SJIS');
        $this->pdf->SetAutoPageBreak(false);
        $this->pdf->SetMargins(0, 0, 0, 0);
        $this->pdf->AddPage('P');
        
    }

    public function writeTitle($title) {
        $this->pdf->SetFont(MINCHO,'', 18);
        $this->pdf->SetXY(16, 16);
        $this->pdf->write(4, s($title));        
    }

    public function writeDate($dateTime) {
        $this->pdf->SetFont(MINCHO,'', 9);
        $this->pdf->SetXY(180, 16);
        $this->pdf->write(4, s(date('Y年m月d日', $dateTime)));
    }

    public function writeCompany($name, $title) {
        $this->pdf->SetFont(MINCHO,'', 16);
        $this->pdf->SetXY(16, 32);
        $this->pdf->write(4, s($name. " ". $title));
    }

    public function writeMyCompany($info) {
        $this->pdf->SetFont(MINCHO,'', 12);
        $this->pdf->SetXY(140, 50);
        $this->pdf->MultiCell(54, 8, s($info), 0, 'L', 0);
    }
    
    public function writeItemTable($x, $y, $items, $taxRate) {
        $this->pdf->SetFont(MINCHO,'', 12);
        $this->pdf->SetXY($x, $y);
        $this->pdf->setFillColor(192, 192, 192);
        $this->pdf->Cell(10, 8, s('No'), 1, 0, 'C', 1);
        $this->pdf->Cell(70, 8, s('品目'), 1, 0, 'C', 1);
        $this->pdf->Cell(20, 8, s('数量'), 1, 0, 'C', 1);
        $this->pdf->Cell(20, 8, s('単位'), 1, 0, 'C', 1);
        $this->pdf->Cell(30, 8, s('単価'), 1, 0, 'C', 1);
        $this->pdf->Cell(30, 8, s('金額'), 1, 1, 'C', 1);
        
        $sum = 0;
        $tax = 0;
        $no = 1;
        foreach ($items as $item) {
            $price = $item['unit_price'] * $item['amount'];
            if ($item['tax_type'] == 1) {
                $tax += getTax($price, $taxRate);
            } else if ($item['tax_type'] == 2) {
                $body = removeTax($price, $taxRate);
                $tax += ($price - $body);
            } 
            
            $sum += $price;
            $this->pdf->SetX(16);
            $this->pdf->Cell(10, 8, s('' + $no), 1, 0, 'C', 0);
            $this->pdf->Cell(70, 8, s($item['subject']), 1, 0, 'L', 0);
            $this->pdf->Cell(20, 8, s($item['amount']), 1, 0, 'R', 0);
            $this->pdf->Cell(20, 8, s($item['degree']), 1, 0, 'L', 0);
            $this->pdf->Cell(30, 8, s('￥'. number_format($item['unit_price'])), 1, 0, 'R', 0);
            $this->pdf->Cell(30, 8, s('￥'. number_format($price)), 1, 1, 'R', 0);
            
            ++$no;
        }
        $total = $sum + $tax;
        
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('小計'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($sum)), 1, 1, 'R', 0);
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('消費税'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($tax)), 1, 1, 'R', 0);
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('合計'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($total)), 1, 1, 'R', 0);
        return array('sum' => $sum, 'tax' => $tax, 'total' => $total);
    }

    public function writeProduct($product) {
        $this->pdf->SetXY(16, $this->pdf->GetY() + 8);
        $this->pdf->Cell(180, 8, s('備考'), 1, 1, 'C', 1);
        $this->pdf->SetX(16);
        $this->pdf->MultiCell(180, 8, s($product), 1, 1, 'L', 0);
    }

    public function writeTotal($total) {
        $this->pdf->SetFont(MINCHO,'', 16);
        $this->pdf->SetXY(16, 100);
        $this->pdf->write(4, s($total));
    }

    public function output($name) {
        $this->pdf->Output($name . '.pdf', 'D');
    }
}
?>